"use strict";

// Seletores do DOM
const display = document.getElementById('display');
const messageArea = document.getElementById('messageArea');
const startBtn = document.getElementById('startBtn');
const pauseBtn = document.getElementById('pauseBtn');
const stopBtn = document.getElementById('stopBtn');

// Estado da Aplicação
let seconds = 0;
let timerInterval = null;
let startTimeISO = null;

/**
 * Converte segundos totais para o formato MM:SS (Minutos acumulados)
 */
const formatTime = (totalSeconds) => {
    const mins = Math.floor(totalSeconds / 60);
    const secs = totalSeconds % 60;
    return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
};

/**
 * Atualiza o display visual
 */
const updateDisplay = () => {
    display.textContent = formatTime(seconds);
};

/**
 * Salva a sessão no LocalStorage
 */
const saveSession = (start, end) => {
    const history = JSON.parse(localStorage.getItem('timer_history') || '[]');
    
    const newEntry = {
        date: new Date().toLocaleDateString('pt-BR'),
        start: start,
        end: end,
        durationSeconds: seconds
    };

    history.push(newEntry);
    localStorage.setItem('timer_history', JSON.stringify(history));
};

// --- Handlers de Evento ---

const startTimer = () => {
    // Limpa mensagem anterior ao iniciar
    messageArea.textContent = "";
    
    // Define o horário de início apenas se for um novo ciclo (seconds === 0)
    if (seconds === 0) {
        startTimeISO = new Date().toLocaleTimeString('pt-BR');
    }

    if (!timerInterval) {
        timerInterval = setInterval(() => {
            seconds++;
            updateDisplay();
        }, 1000);
    }

    // Toggle de botões
    startBtn.disabled = true;
    pauseBtn.disabled = false;
    stopBtn.disabled = false;
};

const pauseTimer = () => {
    clearInterval(timerInterval);
    timerInterval = null;
    
    startBtn.disabled = false;
    pauseBtn.disabled = true;
};

const stopTimer = () => {
    const endTimeISO = new Date().toLocaleTimeString('pt-BR');
    
    // Pausa o intervalo
    pauseTimer();
    
    // Exibe a mensagem final
    const minsFinal = Math.floor(seconds / 60);
    const secsFinal = seconds % 60;
    messageArea.textContent = `${minsFinal} minutos ${secsFinal} segundos`;

    // Persistência
    saveSession(startTimeISO, endTimeISO);

    // Reset de estado interno
    seconds = 0;
    updateDisplay();
    
    // Ajuste de botões
    stopBtn.disabled = true;
};

// Event Listeners
startBtn.addEventListener('click', startTimer);
pauseBtn.addEventListener('click', pauseTimer);
stopBtn.addEventListener('click', stopTimer);
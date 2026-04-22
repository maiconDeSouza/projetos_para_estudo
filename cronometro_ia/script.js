"use strict";

const display = document.getElementById('display');
const messageArea = document.getElementById('messageArea');
const startBtn = document.getElementById('startBtn');
const pauseBtn = document.getElementById('pauseBtn');
const stopBtn = document.getElementById('stopBtn');
const cycleNameInput = document.getElementById('cycleNameInput');
const historyBody = document.getElementById('historyBody');

let seconds = 0;
let timerInterval = null;
let startTimeISO = null;
let currentCycleName = "";

// --- Função de Áudio (Web Audio API) ---
const playNotificationSound = () => {
    const context = new (window.AudioContext || window.webkitAudioContext)();
    const oscillator = context.createOscillator();
    const gain = context.createGain();

    oscillator.type = 'sine';
    oscillator.frequency.setValueAtTime(880, context.currentTime); // Nota Lá
    gain.gain.setValueAtTime(0.1, context.currentTime);

    oscillator.connect(gain);
    gain.connect(context.destination);

    oscillator.start();
    oscillator.stop(context.currentTime + 0.5); // Som por 0.5 segundos
};

const formatTime = (totalSeconds) => {
    const mins = Math.floor(totalSeconds / 60);
    const secs = totalSeconds % 60;
    return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
};

const updateHistoryDisplay = () => {
    const history = JSON.parse(localStorage.getItem('timer_history') || '[]');
    // Pega os últimos 10 e inverte para o mais recente aparecer no topo
    const lastTen = history.slice(-10).reverse();

    historyBody.innerHTML = lastTen.map(item => `
        <tr>
            <td>${item.name || 'Sem nome'}</td>
            <td>${item.date}</td>
            <td>${formatTime(item.durationSeconds)}</td>
        </tr>
    `).join('');
};

const saveSession = (start, end) => {
    const history = JSON.parse(localStorage.getItem('timer_history') || '[]');
    
    const newEntry = {
        name: currentCycleName,
        date: new Date().toLocaleDateString('pt-BR'),
        start: start,
        end: end,
        durationSeconds: seconds
    };

    history.push(newEntry);
    localStorage.setItem('timer_history', JSON.stringify(history));
    updateHistoryDisplay();
};

const startTimer = () => {
    messageArea.textContent = "";
    
    // Se está começando do zero, define o nome e trava o input
    if (seconds === 0) {
        currentCycleName = cycleNameInput.value || "Ciclo s/ Nome";
        startTimeISO = new Date().toLocaleTimeString('pt-BR');
        cycleNameInput.disabled = true;
    }

    if (!timerInterval) {
        timerInterval = setInterval(() => {
            seconds++;
            display.textContent = formatTime(seconds);

            // Regra: Som a cada 30 minutos (1800 segundos)
            if (seconds > 0 && seconds % 1800 === 0) {
                playNotificationSound();
            }
        }, 1000);
    }

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
    pauseTimer();
    
    const minsFinal = Math.floor(seconds / 60);
    const secsFinal = seconds % 60;
    messageArea.textContent = `Tempo Final: ${minsFinal}min ${secsFinal}s`;

    saveSession(startTimeISO, endTimeISO);

    // Reseta para novo ciclo
    seconds = 0;
    display.textContent = "00:00";
    cycleNameInput.disabled = false;
    cycleNameInput.value = "";
    stopBtn.disabled = true;
};

// Event Listeners
startBtn.addEventListener('click', startTimer);
pauseBtn.addEventListener('click', pauseTimer);
stopBtn.addEventListener('click', stopTimer);

// Carrega histórico ao abrir a página
updateHistoryDisplay();
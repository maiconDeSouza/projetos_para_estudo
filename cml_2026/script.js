/**
 * CML 2026 - Lógica do Sistema
 * Desenvolvedor Frontend Sênior
 */

const LANGUAGES = [
    "Assembly", "Fortran", "COBOL", "Lisp", "C", "Swift", "C++", "Java",
    "Python", "JavaScript", "PHP", "C#", "Ruby", "Go", "TypeScript", "Rust"
];

// Estado Inicial do Sistema
let state = {
    status: "setup", // setup, fase1, quartas, semi, final, encerrado
    participantes: [],
    partidas: [], // { round: 1, home: 'C', away: 'Java', score: '' }
    ranking: [],
    playoffs: {
        quartas: [],
        semi: [],
        final: [],
        terceiro: []
    }
};

const STORAGE_KEY = "cml2026";

// --- PERSISTÊNCIA ---

function saveState() {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(state));
    updateRanking();
}

function loadState() {
    const saved = localStorage.getItem(STORAGE_KEY);
    if (saved) {
        state = JSON.parse(saved);
        renderCurrentTab();
    } else {
        resetState();
    }
}

function resetState() {
    state = {
        status: "setup",
        participantes: LANGUAGES.map(name => ({
            name, pts: 0, v: 0, saldo: 0, totalVotos: 0, wins3x0: 0, buchholz: 0, lastRound: 0, bonus: 0
        })),
        partidas: [],
        ranking: [],
        playoffs: { quartas: [], semi: [], final: [], terceiro: [] }
    };
    saveState();
    renderCurrentTab();
}

// --- LÓGICA DE TORNEIO (SORTEIO) ---

function gerarSorteioFase1() {
    const rounds = 5;
    const matches = [];
    const teams = [...LANGUAGES];
    const history = new Set();

    // Algoritmo de emparelhamento simples para evitar repetição
    for (let r = 1; r <= rounds; r++) {
        let pool = [...teams];
        pool.sort(() => Math.random() - 0.5);

        while (pool.length >= 2) {
            const home = pool.shift();
            let awayIndex = pool.findIndex(t => !history.has(`${home}-${t}`) && !history.has(`${t}-${home}`));
            
            if (awayIndex === -1) awayIndex = 0; // Fallback segurança

            const away = pool.splice(awayIndex, 1)[0];
            matches.push({ round: r, home, away, score: "" });
            history.add(`${home}-${away}`);
        }
    }

    state.partidas = matches;
    state.status = "fase1";
    saveState();
    renderCurrentTab();
}

// --- CÁLCULOS DE RANKING ---

function updateRanking() {
    // Reset status
    const stats = {};
    LANGUAGES.forEach(l => {
        stats[l] = { 
            name: l, pts: 0, v: 0, saldo: 0, totalVotos: 0, wins3x0: 0, 
            buchholz: 0, lastRound: 0, opponents: [], bonus: 0 
        };
    });

    // Processar Fase 1
    state.partidas.forEach(m => {
        if (!m.score) return;
        const [v1, v2] = m.score.split('x').map(Number);
        const s1 = stats[m.home];
        const s2 = stats[m.away];

        s1.opponents.push(m.away);
        s2.opponents.push(m.home);
        s1.totalVotos += v1;
        s2.totalVotos += v2;
        s1.saldo += (v1 - v2);
        s2.saldo += (v2 - v1);

        if (v1 > v2) {
            s1.pts += (v1 === 3 ? 3 : 2);
            s2.pts += (v1 === 3 ? 0 : 1);
            s1.v += 1;
            if (v1 === 3) s1.wins3x0 += 1;
        } else {
            s2.pts += (v2 === 3 ? 3 : 2);
            s1.pts += (v2 === 3 ? 0 : 1);
            s2.v += 1;
            if (v2 === 3) s2.wins3x0 += 1;
        }
        if (m.round === 5) {
            s1.lastRound = (v1 > v2 ? v1 : 0);
            s2.lastRound = (v2 > v1 ? v2 : 0);
        }
    });

    // Processar Playoffs (Pontos Acumulativos)
    const playoffMatches = [...state.playoffs.quartas, ...state.playoffs.semi, ...state.playoffs.final, ...state.playoffs.terceiro];
    playoffMatches.forEach(m => {
        if (!m.score) return;
        const [v1, v2] = m.score.split('x').map(Number);
        const s1 = stats[m.home];
        const s2 = stats[m.away];
        
        if (v1 > v2) {
            s1.pts += (v1 === 3 ? 3 : 2);
            s2.pts += (v1 === 3 ? 0 : 1);
        } else {
            s2.pts += (v2 === 3 ? 3 : 2);
            s1.pts += (v2 === 3 ? 0 : 1);
        }
    });

    // Buchholz
    Object.values(stats).forEach(s => {
        s.buchholz = s.opponents.reduce((sum, opp) => sum + stats[opp].pts, 0);
    });

    // Bônus Final
    if (state.status === "encerrado") {
        const f = state.playoffs.final[0];
        const t = state.playoffs.terceiro[0];
        if (f && f.score) {
            const [v1, v2] = f.score.split('x').map(Number);
            stats[v1 > v2 ? f.home : f.away].bonus = 25;
            stats[v1 > v2 ? f.away : f.home].bonus = 20;
        }
        if (t && t.score) {
            const [v1, v2] = t.score.split('x').map(Number);
            stats[v1 > v2 ? t.home : t.away].bonus = 15;
            stats[v1 > v2 ? t.away : t.home].bonus = 10;
        }
    }

    // Ordenação (Critérios de Desempate)
    state.ranking = Object.values(stats).sort((a, b) => {
        const totalA = a.pts + a.bonus;
        const totalB = b.pts + b.bonus;
        if (totalB !== totalA) return totalB - totalA;
        if (b.saldo !== a.saldo) return b.saldo - a.saldo;
        if (b.totalVotos !== a.totalVotos) return b.totalVotos - a.totalVotos;
        if (b.v !== a.v) return b.v - a.v;
        if (b.wins3x0 !== a.wins3x0) return b.wins3x0 - a.wins3x0;
        if (b.buchholz !== a.buchholz) return b.buchholz - a.buchholz;
        return b.lastRound - a.lastRound;
    });
}

// --- RENDERIZAÇÃO ---

let currentTab = "fase1";

function renderCurrentTab() {
    const container = document.getElementById('tab-content');
    updateRanking();

    const tabs = {
        fase1: renderFase1,
        quartas: () => renderPlayoffs("quartas", "Quartas de Final"),
        semi: () => renderPlayoffs("semi", "Semifinais"),
        final: renderFinal,
        regulamento: renderRegulamento
    };

    container.innerHTML = tabs[currentTab]();
    attachListeners();
}

function renderFase1() {
    if (state.status === "setup") {
        return `
            <div class="empty-state">
                <h2>Bem-vindo à CML 2026</h2>
                <p>O torneio ainda não começou. Clique no botão abaixo para realizar o sorteio da primeira fase.</p>
                <br>
                <button onclick="gerarSorteioFase1()" class="btn-primary">Gerar Sorteio da Primeira Fase</button>
            </div>
        `;
    }

    let html = `<div class="dashboard-grid"><div class="rounds-container">`;
    for (let r = 1; r <= 5; r++) {
        html += `<div class="card"><h3 class="card-title">Rodada ${r}</h3>`;
        state.partidas.filter(m => m.round === r).forEach((m, idx) => {
            const globalIdx = state.partidas.indexOf(m);
            html += `
                <div class="match-row">
                    <span class="team-name">${m.home}</span>
                    <select class="score-select" onchange="updateScore('fase1', ${globalIdx}, this.value)">
                        <option value="">-</option>
                        ${["3x0","2x1","1x2","0x3"].map(opt => `<option value="${opt}" ${m.score === opt ? 'selected' : ''}>${opt}</option>`).join('')}
                    </select>
                    <span class="team-name right">${m.away}</span>
                </div>
            `;
        });
        html += `</div>`;
    }
    html += `</div><div>${renderRankingBox(8)}</div></div>`;
    return html;
}

function renderPlayoffs(level, title) {
    const isLocked = !checkPhaseComplete(level === "quartas" ? "fase1" : (level === "semi" ? "quartas" : "semi"));
    
    if (isLocked) return `<div class="empty-state"><h2>Bloqueado</h2><p>Complete a fase anterior para liberar.</p></div>`;

    if (state.playoffs[level].length === 0) {
        return `<div class="empty-state">
            <h2>${title}</h2>
            <button onclick="gerar${level.charAt(0).toUpperCase() + level.slice(1)}()" class="btn-primary">Gerar Confrontos</button>
        </div>`;
    }

    let html = `<div class="dashboard-grid"><div><div class="card"><h3 class="card-title">${title}</h3>`;
    state.playoffs[level].forEach((m, idx) => {
        html += `
            <div class="match-row">
                <span class="team-name">${m.home}</span>
                <select class="score-select" onchange="updateScore('${level}', ${idx}, this.value)">
                    <option value="">-</option>
                    ${["3x0","2x1","1x2","0x3"].map(opt => `<option value="${opt}" ${m.score === opt ? 'selected' : ''}>${opt}</option>`).join('')}
                </select>
                <span class="team-name right">${m.away}</span>
            </div>
        `;
    });
    html += `</div></div><div>${renderRankingBox(level === "quartas" ? 8 : (level === "semi" ? 4 : 16))}</div></div>`;
    return html;
}

function renderFinal() {
    const isLocked = !checkPhaseComplete("semi");
    if (isLocked) return `<div class="empty-state"><h2>Bloqueado</h2><p>As semifinais precisam ser concluídas.</p></div>`;

    if (state.playoffs.final.length === 0) {
        return `<div class="empty-state"><h2>Grande Final</h2><button onclick="gerarFinal()" class="btn-primary">Gerar Final e 3º Lugar</button></div>`;
    }

    const final = state.playoffs.final[0];
    const terceiro = state.playoffs.terceiro[0];

    return `
        <div class="dashboard-grid">
            <div>
                <div class="card">
                    <h3 class="card-title">🏆 Grande Final</h3>
                    <div class="match-row">
                        <span class="team-name">${final.home}</span>
                        <select class="score-select" onchange="updateScore('final', 0, this.value)">
                            <option value="">-</option>
                            ${["3x0","2x1","1x2","0x3"].map(opt => `<option value="${opt}" ${final.score === opt ? 'selected' : ''}>${opt}</option>`).join('')}
                        </select>
                        <span class="team-name right">${final.away}</span>
                    </div>
                </div>
                <div class="card">
                    <h3 class="card-title">🥉 Disputa de 3º Lugar</h3>
                    <div class="match-row">
                        <span class="team-name">${terceiro.home}</span>
                        <select class="score-select" onchange="updateScore('terceiro', 0, this.value)">
                            <option value="">-</option>
                            ${["3x0","2x1","1x2","0x3"].map(opt => `<option value="${opt}" ${terceiro.score === opt ? 'selected' : ''}>${opt}</option>`).join('')}
                        </select>
                        <span class="team-name right">${terceiro.away}</span>
                    </div>
                </div>
            </div>
            <div>${renderRankingBox(16, true)}</div>
        </div>
    `;
}

function renderRankingBox(limit, showAll = false) {
    let rows = showAll ? state.ranking : state.ranking.slice(0, 16);
    let html = `<div class="card">
        <h3 class="card-title">Classificação Geral</h3>
        <table>
            <thead><tr><th>#</th><th>Linguagem</th><th>P</th><th>V</th><th>S</th></tr></thead>
            <tbody>`;
    
    rows.forEach((team, idx) => {
        const isTop = idx < limit;
        const totalPts = team.pts + (team.bonus || 0);
        let medal = "";
        if(state.status === "encerrado") {
            if(idx === 0) medal = "🥇 ";
            else if(idx === 1) medal = "🥈 ";
            else if(idx === 2) medal = "🥉 ";
            else if(idx === 3) medal = "🏅 ";
        }

        html += `<tr class="${isTop ? 'top-8' : ''}">
            <td>${idx + 1}</td>
            <td>${medal}${team.name}</td>
            <td><strong>${totalPts}</strong></td>
            <td>${team.v}</td>
            <td>${team.saldo}</td>
        </tr>`;
    });
    
    html += `</tbody></table></div>`;
    return html;
}

function renderRegulamento() {
    return `
        <div class="card regulamento-text">
            <h2>Regulamento CML 2026</h2>
            <p>A Copa do Mundo das Linguagens é uma competição técnica e de prestígio entre as 16 principais tecnologias de programação.</p>
            <h3>1. Fase de Grupos</h3>
            <ul>
                <li>Cada linguagem joga 5 partidas.</li>
                <li>Pontuação: 3x0 (3 pts), 2x1 (2 pts p/ vencedor, 1 pt p/ perdedor), 0x3 (0 pts).</li>
                <li>Os 8 melhores avançam.</li>
            </ul>
            <h3>2. Critérios de Desempate</h3>
            <ol>
                <li>Pontos (incluindo bônus)</li>
                <li>Saldo de Votos</li>
                <li>Total de Votos Recebidos</li>
                <li>Número de Vitórias</li>
                <li>Buchholz (Força dos oponentes)</li>
            </ol>
            <h3>3. Playoffs</h3>
            <ul>
                <li>Quartas: 1º x 8º, 2º x 7º, 3º x 6º, 4º x 5º.</li>
                <li>Semifinais: Cruzamento baseado na pontuação acumulada (Melhor x Pior).</li>
                <li>Bônus: Campeão (+25), Vice (+20), 3º (+15), 4º (+10).</li>
            </ul>
        </div>
    `;
}

// --- EVENTOS E CONTROLE ---

function updateScore(phase, index, value) {
    if (phase === "fase1") state.partidas[index].score = value;
    else if (phase === "terceiro") state.playoffs.terceiro[index].score = value;
    else state.playoffs[phase][index].score = value;
    
    if (phase === "final") state.status = "encerrado";
    saveState();
    renderCurrentTab();
}

function checkPhaseComplete(phase) {
    if (phase === "fase1") return state.partidas.length > 0 && state.partidas.every(m => m.score !== "");
    return state.playoffs[phase] && state.playoffs[phase].length > 0 && state.playoffs[phase].every(m => m.score !== "");
}

// --- GERAÇÃO DE FASES ---

function gerarQuartas() {
    const top8 = state.ranking.slice(0, 8);
    state.playoffs.quartas = [
        { home: top8[0].name, away: top8[7].name, score: "" },
        { home: top8[1].name, away: top8[6].name, score: "" },
        { home: top8[2].name, away: top8[5].name, score: "" },
        { home: top8[3].name, away: top8[4].name, score: "" }
    ];
    state.status = "quartas";
    saveState();
    renderCurrentTab();
}

function gerarSemi() {
    const winners = [];
    state.playoffs.quartas.forEach(m => {
        const [v1, v2] = m.score.split('x').map(Number);
        winners.push(v1 > v2 ? m.home : m.away);
    });

    // Reordenar vencedores por pontuação acumulada
    updateRanking();
    const sortedWinners = state.ranking.filter(t => winners.includes(t.name));

    state.playoffs.semi = [
        { home: sortedWinners[0].name, away: sortedWinners[3].name, score: "" },
        { home: sortedWinners[1].name, away: sortedWinners[2].name, score: "" }
    ];
    state.status = "semi";
    saveState();
    renderCurrentTab();
}

function gerarFinal() {
    const finalists = [];
    const losers = [];
    state.playoffs.semi.forEach(m => {
        const [v1, v2] = m.score.split('x').map(Number);
        if (v1 > v2) { finalists.push(m.home); losers.push(m.away); }
        else { finalists.push(m.away); losers.push(m.home); }
    });

    state.playoffs.final = [{ home: finalists[0], away: finalists[1], score: "" }];
    state.playoffs.terceiro = [{ home: losers[0], away: losers[1], score: "" }];
    state.status = "final";
    saveState();
    renderCurrentTab();
}

function attachListeners() {
    document.querySelectorAll('.tab-btn').forEach(btn => {
        btn.onclick = () => {
            document.querySelectorAll('.tab-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            currentTab = btn.dataset.tab;
            renderCurrentTab();
        };
    });

    document.getElementById('btn-reset').onclick = () => {
        if (confirm("Deseja realmente resetar todo o torneio?")) resetState();
    };
}

// Inicialização
window.onload = loadState;
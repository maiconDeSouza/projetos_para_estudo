const copaDasLinguagens = [
  "Assembly",
  "Fortran",
  "COBOL",
  "Lisp",
  "C",
  "Swift",
  "C++",
  "Java",
  "Python",
  "JavaScript",
  "PHP",
  "C#",
  "Ruby",
  "Go",
  "TypeScript",
  "Rust"
];

class Copa {
  constructor(linguagens) {
    this.linguagens = linguagens;
    this.rodadas = [];
    this.confrontos = new Set(); // controla jogos já feitos
  }

  // cria chave única independente da ordem
  gerarChave(time1, time2) {
    return [time1, time2].sort().join(" vs ");
  }

  // embaralhamento correto
  embaralhar(array) {
    const copia = [...array];

    for (let i = copia.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [copia[i], copia[j]] = [copia[j], copia[i]];
    }

    return copia;
  }

  gerarRodada() {
    let tentativa = 0;

    while (tentativa < 1000) {
      tentativa++;

      const embaralhado = this.embaralhar(this.linguagens);
      const jogos = [];
      let valido = true;

      for (let i = 0; i < embaralhado.length; i += 2) {
        const time1 = embaralhado[i];
        const time2 = embaralhado[i + 1];

        const chave = this.gerarChave(time1, time2);

        if (this.confrontos.has(chave)) {
          valido = false;
          break;
        }

        jogos.push({ time1, time2, chave });
      }

      if (valido) {
        // salva confrontos
        jogos.forEach(jogo => this.confrontos.add(jogo.chave));

        return jogos.map(j => `${j.time1} vs ${j.time2}`);
      }
    }

    throw new Error("Não foi possível gerar uma rodada válida sem repetir jogos.");
  }

  gerarRodadas(qtd = 5) {
    this.rodadas = [];
    this.confrontos.clear();

    for (let i = 1; i <= qtd; i++) {
      const rodada = {
        nome: `Rodada ${i}`,
        jogos: this.gerarRodada()
      };

      this.rodadas.push(rodada);
    }
  }

  mostrarRodadas() {
    this.rodadas.forEach(rodada => {
      console.log(`\n🏟️ ${rodada.nome}`);
      console.log("----------------------");

      rodada.jogos.forEach(jogo => {
        console.log(jogo);
      });
    });
  }
}

// execução
const copa = new Copa(copaDasLinguagens);

copa.gerarRodadas(5);
copa.mostrarRodadas();
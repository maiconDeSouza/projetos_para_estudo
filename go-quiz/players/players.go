package players

type Players struct {
	nome      string
	pontuacao uint
}

func NovaJogadores(nome string) Players {
	p := Players{nome: nome, pontuacao: 0}
	return p
}

func (p *Players) Acertos() {
	p.pontuacao++
}

func (p Players) PontuacaoTotal() uint {
	return p.pontuacao
}

func (p Players) PegarNome() string {
	return p.nome
}

package main

import (
	"fmt"
	"go-quiz/players"
	"go-quiz/quiz"
	"go-quiz/terminal"
	"strings"
)

func main() {
	terminal.Clear()
	perguntas := quiz.PegarPerguntas()
	fmt.Println("Olá bem vindo ao Quiz GoLang")
	terminal.PrintTrace()
	fmt.Println()

	fmt.Println("Digite seu nome e vamos começar o jogo!")
	nome := terminal.ScanTerminal()

	p := players.NovaJogadores(nome)
	fmt.Printf("Vamos lá %s\n", p.PegarNome())
	terminal.Sleep(3)

	for key, value := range perguntas {
		terminal.Clear()
		fmt.Println("Quiz GoLang")
		terminal.PrintTrace()
		fmt.Println()

		fmt.Printf("Tema da proxima pergunta: %s\n", key)
		terminal.PrintTrace()

		fmt.Println(value.Pergunta)
		r := terminal.ScanTerminal()
		terminal.Sleep(3)

		if strings.EqualFold(r, value.Resposta) {
			terminal.Clear()
			fmt.Println("Certa resposta!!!!")
			terminal.Sleep(2)
			p.Acertos()
			continue
		}
		fmt.Println("Você errou!!!")
		terminal.Sleep(2)
	}
	terminal.Clear()
	fmt.Printf("Você teve %d acertos de %d perguntas", p.PontuacaoTotal(), len(perguntas))
}

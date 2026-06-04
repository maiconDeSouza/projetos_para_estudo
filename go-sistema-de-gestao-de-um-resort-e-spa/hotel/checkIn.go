package hotel

import (
	"bufio"
	"fmt"
	utilsmcn "go-sistema-de-gestao-de-um-resort-e-spa/utils-mcn"
	"os"
	"strings"
)

var PrintMCN = utilsmcn.PrintMCN
var render = bufio.NewReader(os.Stdin)

func CheckOption(op string, nameHotel string) {
	switch op {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
		PrintMCN("Até logo!!!!!")
		utilsmcn.SleepMCN(2)
		utilsmcn.Clear()
	default:
		PrintMCN("Opção inválida!!!!")
		PrintMCN("Tente novamente")
		utilsmcn.SleepMCN(2)
		CheckIn(nameHotel)
	}
}

func CheckIn(nameHotel string) string {
	utilsmcn.Clear()
	for {
		PrintMCN(fmt.Sprintf("Seja bem vindo(a) ao %s!", nameHotel))
		PrintMCN("---------------------")
		PrintMCN("Escolha uma das opções abaixo:")
		PrintMCN(" ")
		PrintMCN("1) Novo hóspede")
		PrintMCN("2) Todos os hóspedes")
		PrintMCN("3) Cadastrar novo serviço")
		PrintMCN("4) Checkout")
		PrintMCN("5) Sair")

		op, err := render.ReadString('\n')
		op = strings.TrimSpace(op)

		if err != nil {
			PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err))
			utilsmcn.Clear()
			continue
		}

		utilsmcn.Clear()
		utilsmcn.SleepMCN(2)
		return op
	}
}

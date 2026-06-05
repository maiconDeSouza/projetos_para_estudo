package terminal

import (
	"bufio"
	"fmt"
	"go-sistema-de-gestao-de-um-resort-e-spa/guests"
	"go-sistema-de-gestao-de-um-resort-e-spa/hotel"
	utilsmcn "go-sistema-de-gestao-de-um-resort-e-spa/utils-mcn"
	"os"
	"strconv"
	"strings"
)

var PrintMCN = utilsmcn.PrintMCN
var render = bufio.NewReader(os.Stdin)
var h hotel.Hotel

func init() {
	h = hotel.NewHotel()
}

func CheckOption(op string) {
	switch op {
	case "1":
		registerGuests()
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
		CheckIn()
	}
}

func CheckIn() {
	utilsmcn.Clear()
	for {
		PrintMCN(fmt.Sprintf("Seja bem vindo(a) ao %s!", h.Name))
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
		CheckOption(op)
		utilsmcn.Clear()
		break

	}
}

func registerGuests() {
	r := h.GetRooms()
	for {
		utilsmcn.Clear()
		PrintMCN("Quartos disponiveis:")
		PrintMCN(" ")

		for i, room := range r {
			PrintMCN(fmt.Sprintf("%d) Quarto: %s", i+1, room.Room))
			PrintMCN(fmt.Sprintf("Descrição do quarto: %s", room.View))
			PrintMCN(fmt.Sprintf("Preço R$%.2f", room.Price))
			PrintMCN(fmt.Sprintf("Ocupado: %t", room.Busy))
			PrintMCN(fmt.Sprintf("Limpo: %t", room.IsItClean))
			PrintMCN("---------------------")
		}

		PrintMCN("Digite qual quarto:")

		op, err := render.ReadString('\n')
		op = strings.TrimSpace(op)
		utilsmcn.Clear()

		if err != nil {
			PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err))
			continue
		}

		n, err1 := strconv.Atoi(op)
		if err1 != nil {
			PrintMCN(fmt.Sprint("Erro na conversão:", err1))
			utilsmcn.SleepMCN(2)
			continue
		}

		if n > len(r) {
			PrintMCN("Número de quarto errado!")
			utilsmcn.SleepMCN(2)
			continue
		}

		if r[n-1].Busy {
			PrintMCN("Quarto ocupado!!!!")
			utilsmcn.SleepMCN(2)
			continue
		}

		PrintMCN("Digite o nome do hóspede:")
		name, err2 := render.ReadString('\n')
		name = strings.TrimSpace(name)

		if err2 != nil {
			PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err2))
			continue
		}

		g := guests.RegisterGuest(name, guests.Premium)
		utilsmcn.SleepMCN(2)
		PrintMCN(fmt.Sprintf("Olá %s você pode cadastrar 4 dependentes. deseja cadastrar?", g.GetName()))
		is, err2 := render.ReadString('\n')

		if is == "n" {
			continue
		}
	}
}

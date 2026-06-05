package terminal

import (
	"bufio"
	"fmt"
	"go-sistema-de-gestao-de-um-resort-e-spa/guests"
	"go-sistema-de-gestao-de-um-resort-e-spa/hotel"
	utilsmcn "go-sistema-de-gestao-de-um-resort-e-spa/utils-mcn"
	"os"
	"strconv"
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
		getTotalGuests()
	case "3":
		newService()
	case "4":
		checkout()
	case "5":
		nextDay()
	case "6":
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

func OptionTier(name string, tier string) guests.Guest {
	switch tier {
	case "1":
		return guests.RegisterGuest(name, guests.Premium)
	case "2":
		return guests.RegisterGuest(name, guests.Intermediate)
	case "3":
		return guests.RegisterGuest(name, guests.Standard)
	default:
		return guests.RegisterGuest(name, guests.Standard)
	}
}

func RegisterDependent(g *guests.Guest) {
	utilsmcn.Clear()
	var err error
	var i int = 1
	for {
		if len(g.GetDependent()) == 4 {
			PrintMCN("Número Máximo de dependentes!!!!")
			utilsmcn.SleepMCN(2)
			break
		}
		PrintMCN(fmt.Sprintf("Cadastre o depedente %d ou deixe em branco para sair", i))
		i++
		opNameDependent := utilsmcn.ReadTerminal(&err)
		if err != nil {
			continue
		}
		if opNameDependent == "" {
			break
		}
		PrintMCN("Digite:")
		PrintMCN("1 para Adulto")
		PrintMCN("2 para criança")
		opIsAdult := utilsmcn.ReadTerminal(&err)
		if err != nil {
			continue
		}

		if opIsAdult == "2" {
			g.AddDependent(opNameDependent, false)
			continue
		}
		g.AddDependent(opNameDependent, true)
	}
}

func CheckIn() {
	utilsmcn.Clear()
	var err error
	for {
		PrintMCN(fmt.Sprintf("Seja bem vindo(a) ao %s!", h.Name))
		PrintMCN("---------------------")
		PrintMCN("Escolha uma das opções abaixo:")
		PrintMCN(" ")
		PrintMCN("1) Novo hóspede")
		PrintMCN("2) Todos os hóspedes")
		PrintMCN("3) Cadastrar novo serviço")
		PrintMCN("4) Checkout")
		PrintMCN("5) Proxímo dia")
		PrintMCN("6) Sair")

		op := utilsmcn.ReadTerminal(&err)

		if err != nil {
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
	var err error
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

		PrintMCN("Digite qual quarto ou N para sair:")

		opRoom := utilsmcn.ReadTerminal(&err)
		utilsmcn.Clear()

		if err != nil {
			continue
		}
		if opRoom == "N" {
			CheckIn()
			break
		}

		n, errNumber := strconv.Atoi(opRoom)
		if errNumber != nil {
			PrintMCN(fmt.Sprint("Erro na conversão:", errNumber))
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
		opName := utilsmcn.ReadTerminal(&err)
		if err != nil {
			continue
		}

		PrintMCN("---------------------")
		PrintMCN("Qual Categoria:")
		PrintMCN("1) Premium")
		PrintMCN("2) Intermediate")
		PrintMCN("3) Standard")
		opTier := utilsmcn.ReadTerminal(&err)
		if err != nil {
			continue
		}

		utilsmcn.Clear()

		g := OptionTier(opName, opTier)
		utilsmcn.SleepMCN(2)

		PrintMCN(fmt.Sprintf("Olá %s você pode cadastrar 4 dependentes. deseja cadastrar?", g.GetName()))
		PrintMCN("Sim")
		PrintMCN("Não")
		opDependents := utilsmcn.ReadTerminal(&err)
		if err != nil {
			continue
		}
		if opDependents == "Sim" {
			RegisterDependent(&g)
		}

		nameRoom := fmt.Sprintf("Quarto: %s - %s", h.GetRooms()[n-1].Room, h.GetRooms()[n-1].View)
		priceRoom := h.GetRooms()[n-1].Price
		h.GetRooms()[n-1].Busy = true

		g.AddService(nameRoom, priceRoom)
		guests.NewGuests(&g)

		utilsmcn.Clear()
		PrintMCN("Quarto cadastrado com sucesso!")
		utilsmcn.SleepMCN(5)
		CheckIn()
		break
	}
}

func getTotalGuests() {
	var err error
	guestsList := guests.GetGuestsList()

	for _, g := range guestsList {
		nameGuest := g.GetName()
		tier := g.GetTier()
		dependent := ""
		services := ""
		day := g.GetDay()

		for i, d := range g.GetDependent() {
			if i == 0 {
				dependent += fmt.Sprintf("%s,", d.Name)
				continue
			}
			if i == 3 {
				dependent += fmt.Sprintf(" %s", d.Name)
				continue
			}
			dependent += fmt.Sprintf(" %s,", d.Name)
		}

		for i, s := range g.GetService() {
			if i == 0 {
				services += fmt.Sprintf("%s - R$%.2f |", s.Name, s.Price)
				continue
			}
			if i == len(g.GetService())-1 {
				services += fmt.Sprintf(" %s - R$%.2f ", s.Name, s.Price)
				continue
			}
			services += fmt.Sprintf(" %s - R$%.2f |", s.Name, s.Price)
		}

		PrintMCN(fmt.Sprintf("Hóspede: %s", nameGuest))
		PrintMCN(fmt.Sprintf("Categoria: %s", tier))
		PrintMCN(fmt.Sprintf("Dependentes: %s", dependent))
		PrintMCN(fmt.Sprintf("Serviços: %s", services))
		PrintMCN(fmt.Sprintf("Dias Hospedados: %d", day))

		PrintMCN("---------------------")

	}
	PrintMCN("Aperte qualquer botão para voltar")
	utilsmcn.ReadTerminal(&err)
	utilsmcn.SleepMCN(3)
	CheckIn()
}

func newService() {
	var err error
	guestsList := guests.GetGuestsList()
	servicesList := h.Services
	PrintMCN("Clientes:")

	for i, g := range guestsList {
		PrintMCN(fmt.Sprintf("%d) %s", i+1, g.GetName()))
		PrintMCN("---------------------")
	}

	PrintMCN("Digite o número do cliente que gostaria de adicionar um serviço:")
	numberClient := utilsmcn.ReadTerminal(&err)
	if err != nil {
		PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err))
	}

	ng, errNumber := strconv.Atoi(numberClient)
	if errNumber != nil {
		PrintMCN(fmt.Sprint("Erro na conversão:", errNumber))
	}
	ng--

	utilsmcn.SleepMCN(2)
	utilsmcn.Clear()

	PrintMCN(fmt.Sprintf("O cliente %s quer qual serviço?", guestsList[ng].GetName()))
	for i, s := range servicesList {
		PrintMCN(fmt.Sprintf("%d) %s", i+1, s.Name))
		PrintMCN("---------------------")
	}

	PrintMCN("Digite o número do serviço:")
	numberService := utilsmcn.ReadTerminal(&err)

	if err != nil {
		PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err))
	}

	ns, errNumber2 := strconv.Atoi(numberService)
	if errNumber2 != nil {
		PrintMCN(fmt.Sprint("Erro na conversão:", errNumber))
	}
	ns--

	utilsmcn.Clear()
	PrintMCN(fmt.Sprintf("O cliente %s está contratando o serviço: %s?", guestsList[ng].GetName(), servicesList[ns].Name))
	PrintMCN("Digite Sim ou Não")
	confirm := utilsmcn.ReadTerminal(&err)

	if err != nil {
		PrintMCN(fmt.Sprint("Erro ao ler o prompt:", err))
	}

	if confirm != "Sim" {
		newService()
	}

	guestsList[ng].AddService(servicesList[ns].Name, servicesList[ns].Price)

	utilsmcn.Clear()
	PrintMCN("Serviço adicionado com sucesso!!!")
	utilsmcn.SleepMCN(3)
	CheckIn()
}

func nextDay() {
	guestsList := guests.GetGuestsList()

	for i := range guestsList {
		guestsList[i].AddDay()
	}

	PrintMCN("Boa Noite!!!")
	utilsmcn.SleepMCN(3)
	CheckIn()
}

func checkout() {
	// var err error
	guestsList := guests.GetGuestsList()

	for i, g := range guestsList {
		total := 0.00
		for i, s := range g.GetService() {
			if i == 0 {
				if g.GetDay() == 0 {
					total += s.Price
					continue
				}
				total += (float64(g.GetDay()) * s.Price)
				continue
			}
			total += s.Price
		}

		total = total - (total * g.GetDiscount())
		PrintMCN(fmt.Sprintf("%d) %s - total: R$%.2f", i+1, g.GetName(), total))
		PrintMCN("---------------------")
	}
	utilsmcn.SleepMCN(25)
}

package main

import (
	"fmt"
	"go-sistema-de-gestao-de-um-resort-e-spa/guests"
	"go-sistema-de-gestao-de-um-resort-e-spa/hotel"
)

func main() {
	h := hotel.NewHotel()

	fmt.Println(h.Services[0].Name)

	g := guests.RegisterGuest("Maicon", guests.Premium)

	fmt.Println(g.GetTier())
	fmt.Println(g.GetDiscount())

}

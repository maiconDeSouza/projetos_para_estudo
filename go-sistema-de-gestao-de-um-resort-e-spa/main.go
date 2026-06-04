package main

import (
	"go-sistema-de-gestao-de-um-resort-e-spa/hotel"
)

func main() {
	h := hotel.NewHotel()

	op := hotel.CheckIn(h.Name)

	hotel.CheckOption(op, h.Name)

	// g := guests.RegisterGuest("Maicon", guests.Premium)

}

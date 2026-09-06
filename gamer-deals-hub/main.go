package main

import (
	"fmt"
	"gamer-deals-hub/database"
)

func main() {
	fmt.Println("🎮 Iniciando Gamer Deals Hub...")

	database.Connect()
}

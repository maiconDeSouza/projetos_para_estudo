package main

import "api-sistema/internal/database"

func main() {
	database.Connect()

	database.CreateAdmin()
}

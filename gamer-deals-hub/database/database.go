package database

import (
	"fmt"

	"gamer-deals-hub/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=secretpassword dbname=gamer_deals port=5454 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar no banco de dados: " + err.Error())
	}

	fmt.Println("⚡ Conexão com o banco de dados PostgreSQL estabelecida!")

	err = db.AutoMigrate(&models.Game{}, &models.Deal{})
	if err != nil {
		panic("Falha ao rodar as migrações: " + err.Error())
	}

	fmt.Println("📊 Tabelas migradas com sucesso!")
	DB = db
}

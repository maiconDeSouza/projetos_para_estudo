package database

import (
	"fmt"
	"log"
	"movies-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Falha ao conectar no banco de dados: %v", err)
	}

	movie := models.Movie{}
	err = movie.AutoMigrate(db)
	if err != nil {
		log.Fatalf("❌ Falha ao migrar movie: %v", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")
	return db
}

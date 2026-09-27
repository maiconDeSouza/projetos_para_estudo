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

	user := models.User{}
	err = user.AutoMigrateUser(db)
	if err != nil {
		log.Fatalf("❌ Falha ao migrar movie: %v", err)
	}

	movie := models.Movie{}
	err = movie.AutoMigrateMovie(db)
	if err != nil {
		log.Fatalf("❌ Falha ao migrar movie: %v", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")
	return db
}

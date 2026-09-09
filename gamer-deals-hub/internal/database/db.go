package database

import (
	"fmt"
	"os"

	"gamer-deals-hub/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect inicializa a conexão via GORM e roda as migrações 🚀
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar via GORM: %w", err)
	}

	// Automigração: Cria ou atualiza as tabelas automaticamente com base nos modelos 🔄
	err = db.AutoMigrate(&models.Game{})
	if err != nil {
		return nil, fmt.Errorf("falha ao executar automigração: %w", err)
	}

	fmt.Println("Conexão com GORM estabelecida e automigração concluída! 🐘")
	return db, nil
}

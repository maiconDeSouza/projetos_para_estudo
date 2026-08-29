package db

import (
	"api-em-tres-cores/internal/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=secretpassword dbname=apitrescores port=1992 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco de dados: %w", err)
	}

	err = database.AutoMigrate(&model.Player{}, &model.Match{}, &model.Event{})
	if err != nil {
		return nil, fmt.Errorf("erro ao executar auto-migration: %w", err)
	}

	return database, nil
}

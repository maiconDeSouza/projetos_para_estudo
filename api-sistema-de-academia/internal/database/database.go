package database

import (
	"api-sistema/internal/config"
	"api-sistema/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d *Database) Connect(cfg *config.ConfigEnv) error {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DB_HOST,
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
		cfg.DB_PORT,
		cfg.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	d.db = db

	err = d.db.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Gym{},
		&models.Visit{},
	)

	if err != nil {
		return err
	}

	log.Println("Banco conectado!")

	return nil
}

func (d *Database) CreateUser(user *models.User) error {
	return d.db.Create(user).Error
}

func NewDatabase(cfg *config.ConfigEnv) (*Database, error) {
	db := &Database{}

	if err := db.Connect(cfg); err != nil {
		return nil, err
	}

	return db, nil
}

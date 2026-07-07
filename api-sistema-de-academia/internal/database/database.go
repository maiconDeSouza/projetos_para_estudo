package database

import (
	"api-sistema/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Gym{},
		&models.Visit{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Banco conectado!")
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func CreateAdmin() {
	var admin models.Admin

	err := DB.Where("nick_name = ?", "admin").First(&admin).Error

	if err == nil {
		return
	}

	password, err := HashPassword("123456")
	if err != nil {
		panic(err)
	}

	admin = models.Admin{
		NickName: "admin",
		Password: password,
	}

	DB.Create(&admin)
}

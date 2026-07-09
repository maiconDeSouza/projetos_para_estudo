package models

import (
	"time"

	"gorm.io/gorm"
)

type Category string

const (
	Basic   Category = "Basico"
	Premium Category = "Prêmio"
)

type Admin struct {
	gorm.Model
	NickName string
	Password string
}

type User struct {
	gorm.Model
	NickName string
	Password string
	Category Category `gorm:"type:varchar(20)"`
}

type UserRequest struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Category string `json:"category"`
}

type Gym struct {
	gorm.Model
	Name     string
	Category Category `gorm:"type:varchar(20)"`
}

type Visit struct {
	gorm.Model

	UserID uint
	User   User

	GymID uint
	Gym   Gym

	CheckIn  time.Time
	CheckOut time.Time
}

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

type ADM struct {
	gorm.Model
	NickName string
	Password string
}

type User struct {
	gorm.Model
	NickName string
	Password string
	Category Category
}

type Gym struct {
	gorm.Model
	Name     string
	Category Category
}

type CheckInCheckOut struct {
	ID       uint
	User     User
	Gym      Gym
	Checkin  time.Time
	checkout time.Time
}

package model

import (
	"uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;uniqueIndex;not null;column:id"`
	Name     string    `json:"name" gorm:"type:varchar(255);not null;column:name"`
	Password string    `json:"password" gorm:"type:varchar(255);not null;column:password"`
	Games    []Game    `json:"games" gorm:"foreignKey:UserID"`
}

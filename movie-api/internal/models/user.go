package models

import (
	"uuid"

	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary" json:"id"`
	Email    string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string    `json:"-"`
	Movie    []Movie   `gorm:"many2many:user_movies;" json:"movies,omitempty"`
}

func (u User) AutoMigrateUser(db *gorm.DB) error {
	return db.AutoMigrate(&u)
}

type NewUserRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenJTWResponse struct {
	Token string `json:"token"`
}

type UserMovie struct {
	User_ID  uuid.UUID `gorm:"column:user_id;primaryKey"`
	Movie_ID uuid.UUID `gorm:"column:movie_id;primaryKey"`
}

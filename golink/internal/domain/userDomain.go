package domain

import (
	"time"
	"uuid"
)

type Link struct {
	ID        uuid.UUID `bson:"_id"`
	Title     string    `bson:"title"`
	Link      string    `bson:"link"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"update_at"`
	Count     uint      `bson:"count"`
}

type User struct {
	ID       uuid.UUID `bson:"_id"`
	Name     string    `bson:"name"`
	Age      uint8     `bson:"age"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Admin    bool      `bson:"admin"`
	Links    []Link    `bson:"links"`
}

type UserRequest struct {
	Name           string `json:"name" validate:"required,min=4,max=25"`
	Email          string `json:"email" validate:"required,email"`
	Age            uint8  `json:"age" validate:"required,min=1,max=140"`
	Password       string `json:"password" validate:"required,min=4,max=55"`
	RepeatPassword string `json:"repeatPassword" validate:"required,min=4,max=55"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=55"`
}

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Admin bool      `json:"admin"`
	Links []Link    `json:"links"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Age   uint8     `json:"age"`
	Admin bool      `json:"admin"`
	Links []Link    `json:"links"`
}

type LinkRequest struct {
	Title string `json:"title" validate:"required,min=4,max=25"`
	Link  string `json:"link" validate:"required,url"`
}

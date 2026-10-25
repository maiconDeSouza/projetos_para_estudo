package model

type User struct {
	Name     string `bson:"name" json:"name" validate:"required,min=3,max=25"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required,min=3,max=25"`
}

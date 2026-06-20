package models

type Book struct {
	ID      string           `json:"id"`
	Name    string           `json:"name"`
	Author  string           `json:"author"`
	History map[string]*User `json:"history"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

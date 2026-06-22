package models

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Borrowed bool   `json:"borrowed"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

package models

type BookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type BookResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Borrowed bool   `json:"borrowed"`
}

type UserRequest struct {
	Name string `json:"name"`
}

type UserResponse struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Books []string `json:"books"`
}

type BorrowedResponse struct {
	NameUser string
	NameBook string
}

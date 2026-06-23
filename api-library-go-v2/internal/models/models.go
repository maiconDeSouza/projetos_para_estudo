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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	ID    int            `json:"id"`
	Name  string         `json:"name"`
	Books []*BookRequest `json:"books"`
}

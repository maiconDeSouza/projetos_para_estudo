package models

type Product struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount uint    `json:"amount"`
}

type ProductResquest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Inventory struct {
	Amount uint `json:"amount"`
}

type OrderRequest struct {
	ID         uint   `json:"id"`
	ClientName string `json:"clienteName"`
	IdProduct  uint   `json:"idProduct"`
}

package repositories

import (
	"api-ecommerce/internal/models"
	"errors"
	"sync"
)

var gate sync.RWMutex

type ResotiroriesInterface interface {
	GetProducts() map[uint]*models.Product
	GetProduct(id uint) (*models.Product, error)
}

type DB struct {
	Products map[uint]*models.Product
	Sales    []*models.OrderRequest
}

func (db *DB) GetProducts() map[uint]*models.Product {
	gate.RLock()
	defer gate.RUnlock()
	return db.Products
}

func (db *DB) GetProduct(id uint) (*models.Product, error) {
	prod, ok := db.Products[id]
	if !ok {
		return nil, errors.New("Produto não consta no sistema")
	}

	return prod, nil
}

func NewRepositorie() *DB {
	p1 := models.Product{
		ID:     1,
		Name:   "Caneta",
		Price:  3.23,
		Amount: 3,
	}

	p2 := models.Product{
		ID:     2,
		Name:   "Lápis",
		Price:  1.23,
		Amount: 2,
	}

	p3 := models.Product{
		ID:     3,
		Name:   "Borracha",
		Price:  0.57,
		Amount: 1,
	}

	db := &DB{
		Products: make(map[uint]*models.Product),
		Sales:    make([]*models.OrderRequest, 0),
	}

	db.Products[p1.ID] = &p1
	db.Products[p2.ID] = &p2
	db.Products[p3.ID] = &p3

	return db
}

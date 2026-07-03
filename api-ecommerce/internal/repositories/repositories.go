package repositories

import (
	"api-ecommerce/internal/models"
	"errors"
	"fmt"
	"sync"
)

var gate sync.RWMutex

type ResotiroriesInterface interface {
	GetProducts() map[uint]*models.Product
	GetProduct(id uint) (*models.Product, error)
	UpdateProductName(prod *models.Product, newName string) *models.Product
	UpdateProductPrice(prod *models.Product, newPrice float64) *models.Product
	AddAmount(id uint, newAmount uint) (*models.Product, error)
	DecreaseAmout(id uint, amount uint) (*models.Product, error)
	Order(orderRequest *models.OrderRequest) (*models.Order, error)
}

type DB struct {
	Products map[uint]*models.Product
	Sales    []*models.Order
}

func (db *DB) GetProducts() map[uint]*models.Product {
	gate.RLock()
	defer gate.RUnlock()
	return db.Products
}

func (db *DB) GetProduct(id uint) (*models.Product, error) {
	gate.RLock()
	defer gate.RUnlock()
	prod, ok := db.Products[id]
	if !ok {
		return nil, errors.New("Produto não consta no sistema")
	}

	return prod, nil
}

func (db *DB) UpdateProductName(prod *models.Product, newName string) *models.Product {
	gate.Lock()
	defer gate.Unlock()

	prod.Name = newName

	return prod
}

func (db *DB) UpdateProductPrice(prod *models.Product, newPrice float64) *models.Product {
	gate.Lock()
	defer gate.Unlock()

	prod.Price = newPrice

	return prod
}

func (db *DB) AddAmount(id uint, newAmount uint) (*models.Product, error) {
	gate.Lock()
	defer gate.Unlock()

	prod, err := db.GetProduct(id)
	if err != nil {
		return nil, err
	}

	if newAmount <= 0 {
		return nil, errors.New("Não pode adicionar a quantidade 0")
	}

	prod.Amount += newAmount

	return prod, nil
}

func (db *DB) DecreaseAmout(id uint, amount uint) (*models.Product, error) {
	gate.Lock()
	defer gate.Unlock()

	prod, err := db.GetProduct(id)
	if err != nil {
		return nil, err
	}

	if amount > prod.Amount {
		str := fmt.Sprintf("Temos apenas %d em estoque", prod.Amount)
		return nil, errors.New(str)
	}

	prod.Amount -= amount

	return prod, nil
}

func (db *DB) Order(orderRequest *models.OrderRequest) (*models.Order, error) {
	gate.Lock()
	defer gate.Unlock()

	or := models.Order{
		ID:         uint(len(db.Sales) + 1),
		ClientName: orderRequest.ClientName,
		IdProduct:  orderRequest.IdProduct,
		Amount:     orderRequest.Amount,
	}

	_, err := db.DecreaseAmout(or.ID, or.Amount)
	if err != nil {
		return nil, err
	}

	db.Sales = append(db.Sales, &or)

	return &or, nil
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
		Sales:    make([]*models.Order, 0),
	}

	db.Products[p1.ID] = &p1
	db.Products[p2.ID] = &p2
	db.Products[p3.ID] = &p3

	return db
}

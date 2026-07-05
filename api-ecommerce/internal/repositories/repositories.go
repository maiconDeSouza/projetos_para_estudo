package repositories

import (
	"api-ecommerce/internal/models"
	"errors"
	"fmt"
	"sync"
)

var gate sync.RWMutex

type RepositoriesInterface interface {
	GetProducts() map[uint]*models.Product
	GetProduct(id uint) (*models.Product, error)
	UpdateProduct(id uint, newProduct *models.ProductResquest) (*models.Product, error)
	AddAmount(id uint, newAmount uint) (*models.Product, error)
	DecreaseAmount(id uint, amount uint) (*models.Product, error)
	Order(orderRequest *models.OrderRequest) (*models.Order, error)
	GetSales() []*models.Order
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
		return nil, errors.New("produto não consta no sistema")
	}

	return prod, nil
}

<<<<<<< HEAD
func (db *DB) UpdateProduct(id uint, newProduct *models.ProductResquest) (*models.Product, error) {
=======
func (db *DB) GetSales() []*models.Order {
	gate.RLock()
	defer gate.RUnlock()

	return db.Sales
}

func (db *DB) UpdateProductName(prod *models.Product, newName string) *models.Product {
>>>>>>> bcd6169755fd3441794c2cc943c11e062fe70d9c
	gate.Lock()
	defer gate.Unlock()

	prod, ok := db.Products[id]
	if !ok {
		return nil, errors.New("Produto não consta no sistema")
	}

	if newProduct.Name != "" && newProduct.Name != prod.Name {
		prod.Name = newProduct.Name
	}

	if newProduct.Price != prod.Price && newProduct.Price > 0 {
		prod.Price = newProduct.Price
	}

	return prod, nil
}

func (db *DB) AddAmount(id uint, newAmount uint) (*models.Product, error) {
	gate.Lock()
	defer gate.Unlock()

	prod, ok := db.Products[id]
	if !ok {
<<<<<<< HEAD
		return nil, errors.New("Produto não consta no sistema")
=======
		return nil, errors.New("produto não consta no sistema")
>>>>>>> bcd6169755fd3441794c2cc943c11e062fe70d9c
	}

	if newAmount == 0 {
		return nil, errors.New("não pode adicionar quantidade 0")
	}

	prod.Amount += newAmount

	return prod, nil
}

<<<<<<< HEAD
func (db *DB) DecreaseAmout(id uint, amount uint) (*models.Product, error) {
	prod, ok := db.Products[id]
	if !ok {
		return nil, errors.New("Produto não consta no sistema")
=======
func (db *DB) DecreaseAmount(id uint, amount uint) (*models.Product, error) {
	gate.Lock()
	defer gate.Unlock()

	prod, ok := db.Products[id]
	if !ok {
		return nil, errors.New("produto não consta no sistema")
	}

	if amount == 0 {
		return nil, errors.New("a quantidade deve ser maior que 0")
>>>>>>> bcd6169755fd3441794c2cc943c11e062fe70d9c
	}

	if amount > prod.Amount {
		return nil, fmt.Errorf("temos apenas %d em estoque", prod.Amount)
	}

	prod.Amount -= amount

	return prod, nil
}

func (db *DB) Order(orderRequest *models.OrderRequest) (*models.Order, error) {
	gate.Lock()
	defer gate.Unlock()

	prod, ok := db.Products[orderRequest.IdProduct]
	if !ok {
		return nil, errors.New("produto não consta no sistema")
	}

	if orderRequest.Amount == 0 {
		return nil, errors.New("a quantidade deve ser maior que 0")
	}

	if orderRequest.Amount > prod.Amount {
		return nil, fmt.Errorf("temos apenas %d em estoque", prod.Amount)
	}

	prod.Amount -= orderRequest.Amount

	order := &models.Order{
		ID:         uint(len(db.Sales) + 1),
		ClientName: orderRequest.ClientName,
		IdProduct:  orderRequest.IdProduct,
		Amount:     orderRequest.Amount,
	}

	db.Sales = append(db.Sales, order)

	return order, nil
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

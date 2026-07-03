package services

import (
	"api-ecommerce/internal/models"
	"api-ecommerce/internal/repositories"
	"errors"
	"strconv"
)

type ServicesInterface interface {
	GetAllProducts() []*models.Product
	UpProduct(id string, newProduct *models.ProductResquest) (*models.Product, error)
	GetProduct(idString string) (*models.Product, error)
	AddItem(idString string, amount uint) (*models.Product, error)
}

type Services struct {
	repo repositories.ResotiroriesInterface
}

func (s *Services) GetAllProducts() []*models.Product {
	listProd := []*models.Product{}
	prod := s.repo.GetProducts()

	for _, v := range prod {
		listProd = append(listProd, v)
	}

	return listProd

}

func (s *Services) GetProduct(idString string) (*models.Product, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Apenas id númericos")
	}

	prod, err := s.repo.GetProduct(uint(id))
	if err != nil {
		return nil, err
	}

	return prod, nil
}

func (s *Services) UpProduct(idString string, newProduct *models.ProductResquest) (*models.Product, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Apenas id númericos")
	}
	prod, err := s.repo.GetProduct(uint(id))
	if err != nil {
		return nil, err
	}
	if newProduct.Name != "" {
		prod.Name = newProduct.Name
	}

	if newProduct.Price != prod.Price && newProduct.Price > 0 {
		prod.Price = newProduct.Price
	}

	return prod, nil
}

func (s *Services) AddItem(idString string, amount uint) (*models.Product, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Apenas id númericos")
	}

	if amount <= 0 {
		return nil, errors.New("Não pode adicionar a quantidade 0")
	}

	prod, err := s.repo.GetProduct(uint(id))
	if err != nil {
		return nil, err
	}

	prod.Amount += amount

	return prod, nil
}

func NewServices(repo repositories.ResotiroriesInterface) *Services {
	services := &Services{
		repo: repo,
	}
	return services
}

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
	Order(orderRequest *models.OrderRequest) (*models.Order, error)
	GetSales() []*models.Order
}

type Services struct {
	repo repositories.RepositoriesInterface
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

func (s *Services) GetSales() []*models.Order {
	return s.repo.GetSales()
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
	if newProduct.Name != "" && newProduct.Name != prod.Name {
		prod = s.repo.UpdateProductName(prod, newProduct.Name)
	}

	if newProduct.Price != prod.Price && newProduct.Price > 0 {
		prod = s.repo.UpdateProductPrice(prod, newProduct.Price)
	}

	return prod, nil
}

func (s *Services) AddItem(idString string, amount uint) (*models.Product, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Apenas id númericos")
	}

	prod, err := s.repo.AddAmount(uint(id), amount)
	if err != nil {
		return nil, err
	}

	return prod, nil

}

func (s *Services) Order(orderRequest *models.OrderRequest) (*models.Order, error) {
	or, err := s.repo.Order(orderRequest)
	if err != nil {
		return nil, err
	}

	return or, nil
}

func NewServices(repo repositories.RepositoriesInterface) *Services {
	services := &Services{
		repo: repo,
	}
	return services
}

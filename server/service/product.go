package service

import (
	"server/models"
	"server/repository"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetProducts() ([]models.Product, error) {
	return s.repo.GetProducts()
}

func (s *productService) GetProductById(id int64) (models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(product models.Product) (models.Product, error) {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(product models.Product) (models.Product, error) {
	return s.repo.DeleteProduct(product)
}

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

package service

import "server/models"

// Product Service
type ProductService interface {
	GetProducts() ([]models.Product, error)
	GetProductById(int64) (models.Product, error)
	CreateProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
	DeleteProduct(models.Product) (models.Product, error)
}

// Employee Service
type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	CreateEmployee(models.Employee) (bool, error)
}

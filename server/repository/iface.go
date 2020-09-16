package repository

import (
	"server/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductById(int64) (models.Product, error)
	CreateProduct(models.Product) (models.Product, error)
	UpdateProduct(models.Product) (models.Product, error)
	DeleteProduct(models.Product) (models.Product, error)
}

type EmployeeRepository interface {
	GetEmployees() ([]models.Employee, error)
	CreateEmployee(models.Employee) (bool, error)
}

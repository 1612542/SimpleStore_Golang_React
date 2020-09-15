package service

import "server/models"

// Product Service
type ProductService interface {
	GetProducts() ([]models.Product, error)
}

// Employee Service
type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	CreateEmployee(models.Employee) (bool, error)
}

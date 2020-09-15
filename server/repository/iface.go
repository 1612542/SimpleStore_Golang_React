package repository

import (
	"server/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
}

type EmployeeRepository interface {
	GetEmployees() ([]models.Employee, error)
	CreateEmployee(models.Employee) (bool, error)
}

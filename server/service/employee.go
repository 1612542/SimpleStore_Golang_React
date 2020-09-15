package service

import (
	"server/models"
	"server/repository"
)

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s employeeService) GetEmployees() ([]models.Employee, error) {
	return s.repo.GetEmployees()
}

func (s employeeService) CreateEmployee(employee models.Employee) (bool, error) {
	return s.repo.CreateEmployee(employee)
}

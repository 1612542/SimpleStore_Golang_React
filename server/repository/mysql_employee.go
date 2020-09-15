package repository

import (
	"log"
	"server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlEmployeeRepository struct {
	DbConnection *gorm.DB
}

func NewMysqlEmployeeRepository(DbConn *gorm.DB) EmployeeRepository {
	return &mysqlEmployeeRepository{DbConn}
}

func (m *mysqlEmployeeRepository) GetEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	result := m.DbConnection.Table("employee").Find(&employees)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return employees, nil
}

func (m *mysqlEmployeeRepository) CreateEmployee(employee models.Employee) (bool, error) {
	if err := m.DbConnection.Table("employee").Create(&employee).Error; err != nil {
		return false, err
	}
	return true, nil
}

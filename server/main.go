package main

import (
	"encoding/json"
	"fmt"
	"log"
	"server/config"
	"server/connector"
	"server/models"
	"server/repository"
	"server/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open up db connection
	db := connector.ConnectDB(config.MySQLConfig())

	productRepo := repository.NewMysqlProductRepository(db)
	productService := service.NewProductService(productRepo)
	products, err := productService.GetProducts()
	if err != nil {
		log.Fatal(err)
	}
	j, err := json.Marshal(products)
	fmt.Println(string(j))

	employeeRepo := repository.NewMysqlEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employee1 := models.Employee{EmployeeId: "E2", EmployeeName: "employee2"}
	_, err1 := employeeService.CreateEmployee(employee1)
	if err1 != nil {
		log.Fatal(err1)
	}

	employees, err := employeeService.GetEmployees()
	k, err := json.Marshal(employees)
	fmt.Println(string(k))

	employee2 := models.Employee{EmployeeId: "E3", EmployeeName: "employee3"}
	_, err2 := employeeService.CreateEmployee(employee2)
	if err2 != nil {
		log.Fatal(err2)
	}

	employees1, err := employeeService.GetEmployees()
	k1, err := json.Marshal(employees1)
	fmt.Println(string(k1))
}

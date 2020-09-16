package main

import (
	"server/api"
	"server/config"
	"server/connector"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	// Open up db connection
	db := connector.ConnectDB(config.MySQLConfig())
	// db := connector.ConnectDB()
	// e := echo.New()

	// productRepo := repository.NewMysqlProductRepository(db)
	// productService := service.NewProductService(productRepo)
	// products, err := productService.GetProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// j, err := json.Marshal(products)
	// fmt.Println(string(j))

	// employeeRepo := repository.NewMysqlEmployeeRepository(db)
	// employeeService := service.NewEmployeeService(employeeRepo)
	// employee1 := models.Employee{EmployeeId: "E1", EmployeeName: "employee1"}
	// _, err := employeeService.CreateEmployee(employee1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// employees, err := employeeService.GetEmployees()
	// k, err := json.Marshal(employees)
	// fmt.Println(string(k))

	// employee2 := models.Employee{EmployeeId: "E3", EmployeeName: "employee3"}
	// _, err2 := employeeService.CreateEmployee(employee2)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	// employees1, err := employeeService.GetEmployees()
	// k1, err := json.Marshal(employees1)
	// fmt.Println(string(k1))

	// e := echo.New()
	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// //CORS
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	// }))
	// // Server
	// e.Run(standard.New(":1323"))

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	}))
	// api.HandleControllers(e, db)
	var controllers = api.NewControllers(db)
	controllers.MakeHandlers(e)
	e.Logger.Fatal(e.Start(":1323"))
}

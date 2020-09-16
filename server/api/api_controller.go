package api

import (
	"net/http"
	"server/repository"
	"server/service"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// var controllers = map[string]func(e *echo.Echo, db *gorm.DB){
// 	"product": MakeProductHandler,
// }

// func HandleControllers(e *echo.Echo, db *gorm.DB) {
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	for k := range controllers {
// 		controllers[k](e, db)
// 	}
// }

type controllers struct {
	db *gorm.DB
}

func NewControllers(db *gorm.DB) *controllers {
	return &controllers{db}
}

func (c *controllers) MakeHandlers(e *echo.Echo) {
	e.GET("/", IndexHandler)

	productRepo := repository.NewMysqlProductRepository(c.db)
	productService := service.NewProductService(productRepo)
	ProductController := NewProductController(productService)
	ProductController.MakeProductHandler(e)

	// MakeEmployeeHandler(e, c.db)
}

func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

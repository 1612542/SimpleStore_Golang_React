package api

import (
	"net/http"
	"server/models"
	"server/service"
	"strconv"

	"github.com/labstack/echo"
)

type productController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) *productController {
	return &productController{s}
}

func (p *productController) getProducts(c echo.Context) error {
	products, err := p.service.GetProducts()
	if err != nil {
		return echo.NewHTTPError(403, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (p *productController) getProductById(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	product, err := p.service.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(403, err)
	}
	return c.JSON(http.StatusOK, product)
}

func (p *productController) createProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	createdProduct, err := p.service.CreateProduct(product)
	if err != nil {
		return echo.NewHTTPError(403, err)
	}
	return c.JSON(http.StatusOK, createdProduct)
}

func (p *productController) updateProduct(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(403, err)
	}
	product, err := p.service.GetProductById(id)
	if err != nil {
		return echo.NewHTTPError(403, err)
	}
	request := models.Product{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	product.Id = request.Id
	product.ProductName = request.ProductName
	product.Price = request.Price
	product.Quantity = request.Quantity
	updatedProduct, err := p.service.UpdateProduct(product)
	if err != nil {
		return echo.NewHTTPError(403, err)
	}
	return c.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) MakeProductHandler(e *echo.Echo) {
	e.GET("/product", p.getProducts)
	e.GET("/product/:id", p.getProductById)
	e.POST("/product", p.createProduct)
	e.PUT("/product/:id", p.updateProduct)
}

// func MakeProductHandler(e *echo.Echo, db *gorm.DB) {
// 	productRepo := repository.NewMysqlProductRepository(db)
// 	productService := service.NewProductService(productRepo)

// 	e.GET("/product", func(c echo.Context) error {
// 		products, err := productService.GetProducts()
// 		if err != nil {
// 			return echo.NewHTTPError(403, err)
// 		}
// 		return c.JSON(http.StatusOK, products)
// 	})

// 	e.GET("/product/:id", func(c echo.Context) error {
// 		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
// 		product, err := productService.GetProductById(id)
// 		if err != nil {
// 			return echo.NewHTTPError(403, err)
// 		}
// 		return c.JSON(http.StatusOK, product)
// 	})

// 	e.POST("/product", func(c echo.Context) error {
// 		var product models.Product
// 		if err := c.Bind(&product); err != nil {
// 			return err
// 		}
// 		createdProduct, err := productService.CreateProduct(product)
// 		if err != nil {
// 			return echo.NewHTTPError(403, err)
// 		}
// 		return c.JSON(http.StatusOK, createdProduct)
// 	})

// 	e.PUT("/product/:id", func(c echo.Context) error {
// 		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 		fmt.Println(id)
// 		if err != nil {
// 			return echo.NewHTTPError(403, err)
// 		}
// 		product, err := productService.GetProductById(id)
// 		if err != nil {
// 			return echo.NewHTTPError(403, err)
// 		}
// 		request := models.Product{}
// 		if err := c.Bind(&request); err != nil {
// 			return err
// 		}
// 		product = request
// 		updatedProduct, err := productService.UpdateProduct(product)
// 		if err != nil {
// 			return echo.NewHTTPError(403, err)
// 		}
// 		return c.JSON(http.StatusOK, updatedProduct)
// 	})
// }

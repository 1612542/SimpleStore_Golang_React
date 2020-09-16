// package repository

// import (
// 	"database/sql"

// 	"server/models"
// )

// type mysqlProductRepository struct {
// 	DbConnection *sql.DB
// }

// func NewMysqlProductRepository(DbConn *sql.DB) ProductRepository {
// 	return &mysqlProductRepository{DbConn}
// }

// func (m *mysqlProductRepository) GetProducts() ([]models.Product, error) {
// 	result, err := m.DbConnection.Query("select * from product")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer result.Close()

// 	var products []models.Product
// 	for result.Next() {
// 		product := models.Product{}
// 		err := result.Scan(&product.Id, &product.ProductName, &product.Quantity, &product.Price, &product.CreatedAt, &product.UpdatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, product)
// 	}
// 	defer result.Close()

// 	return products, nil
// }

package repository

import (
	"server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlProductRepository struct {
	DbConnection *gorm.DB
}

func NewMysqlProductRepository(DbConn *gorm.DB) ProductRepository {
	return &mysqlProductRepository{DbConn}
}

func (m *mysqlProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := m.DbConnection.Table("product").Find(&products)
	return products, result.Error
}

func (m *mysqlProductRepository) GetProductById(id int64) (models.Product, error) {
	var product models.Product
	result := m.DbConnection.Table("product").First(&product, id)
	return product, result.Error
}

func (m *mysqlProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	result := m.DbConnection.Table("product").Create(&product)
	return product, result.Error
}

func (m *mysqlProductRepository) UpdateProduct(product models.Product) (models.Product, error) {
	result := m.DbConnection.Table("product").Save(&product)
	return product, result.Error
}

func (m *mysqlProductRepository) DeleteProduct(product models.Product) (models.Product, error) {
	result := m.DbConnection.Table("product").Delete(&product)
	return product, result.Error
}

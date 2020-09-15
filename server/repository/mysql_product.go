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
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

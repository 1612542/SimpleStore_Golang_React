package models

import "time"

// Product model
type Product struct {
	Id          int64     `json:"id"`
	ProductCode string    `json:"product_code"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// --------- Table Name --------- //
func (Product) TableName() string {
	return "product"
}

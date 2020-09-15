// package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

// type TodoList struct {
// 	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Task   string             `json:"task,omitempty"`
// 	Status bool               `json:"status,omitempty"`
// }

package models

import "time"

// Product model
type Product struct {
	Id          int64     `json:"id"`
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

package models

import "time"

type Employee struct {
	Id           int64     `json:"id"`
	EmployeeId   string    `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// --------- Table Name --------- //
func (Employee) TableName() string {
	return "employee"
}

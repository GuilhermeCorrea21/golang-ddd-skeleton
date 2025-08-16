package customer_dto

import "gorm.io/gorm"

type CreateCustomerStruct struct {
	Name   string `json:"name" binding:"required"`
	Active bool   `json:"active"`
	gorm.Model
}

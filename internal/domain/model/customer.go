package model

import "gorm.io/gorm"

type Customer struct {
	ID     uint   `gorm:"primaryKey,autoIncrement"`
	Name   string `gorm:"type:varchar(191)"`
	Active bool   `gorm:"default:true"`
	gorm.Model
}

type CustomerJoin struct {
	CustomerID uint      `gorm:"column:customer_id;not_null"`
	Customer   *Customer `gorm:"foreignKey:customer_id;references:id;"`
}

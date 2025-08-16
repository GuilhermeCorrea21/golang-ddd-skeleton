package model

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	TotalAmount int       `gorm:"type:int;not null"`
	DueDate     time.Time `gorm:"not null"`
	CustomerJoin
	gorm.Model
}

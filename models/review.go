package models

import (
	"time"
)

type Review struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint
	Product    Product
	CustomerID uint
	Customer   Customer
	Rating     int    // Rating out of 5
	Comment    string // Review text
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

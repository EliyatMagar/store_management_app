package models

import (
	"time"
)

type Review struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint
	Product    Product `gorm:"foreignKey:ProductID" json:"-" binding:"-"`
	CustomerID uint
	Customer   Customer `gorm:"foreignKey:CustomerID" json:"-" binding:"-"`
	Rating     int
	Comment    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

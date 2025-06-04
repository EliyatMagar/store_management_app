package models

import (
	"time"
)

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	SKU         string

	CategoryID   uint           `json:"category_id" binding:"required"`
	Category     Category       `gorm:"foreignKey:CategoryID" json:"category"`
	InvenoryLogs []InventoryLog `gorm:"foreignKey:ProductID" json:"inventory_logs"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

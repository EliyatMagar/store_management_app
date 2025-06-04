package models

import "time"

type InventoryLog struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`

	QuantityChange int       `json:"quantity_change"`
	Type           string    `json:"type"` // addition, removal, adjustment
	Date           time.Time `json:"date"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`
}

// Input struct for controller: ProductName to lookup, UserID provided directly
type InventoryLogInput struct {
	ProductName    string    `json:"product_name" binding:"required"` // lookup product by name
	UserID         uint      `json:"user_id" binding:"required"`      // user ID provided directly
	QuantityChange int       `json:"quantity_change" binding:"required"`
	Type           string    `json:"type" binding:"required"` // addition, removal, adjustment
	Date           time.Time `json:"date"`                    // optional, will be overridden
}

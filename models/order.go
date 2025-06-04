package models

import "time"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Customer   Customer

	UserID uint
	User   User

	OrderDate   time.Time
	TotalAmount float64
	Status      string `gorm:"default:'pending'"` // pending, completed, cancelled

	OrderItems      []OrderItem
	ShippingAddress string
	PaymentMethod   string `gorm:"default:'credit_card'"` // credit_card, paypal, bank_transfer

}

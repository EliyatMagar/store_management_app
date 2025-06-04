package models

type OrderItem struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint
	Order   Order

	ProductID uint
	Product   Product

	Quantity        int
	PriceAtPurchase float64 // Price at the time of purchase
	Discount        float64 // Any discount applied to this item

}

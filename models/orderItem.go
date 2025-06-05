package models

type OrderItem struct {
	ID      uint   `gorm:"primaryKey"`
	OrderID uint   `json:"orderID"`
	Order   *Order `gorm:"foreignKey:OrderID" json:"-" binding:"-"`

	ProductID uint     `json:"productID"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"-" binding:"-"`

	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"priceAtPurchase"`
	Discount        float64 `json:"discount"`
}

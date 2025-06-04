package models

type Customer struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	Orders  []Order  `json:"orders,omitempty"`
	Reviews []Review `json:"reviews,omitempty"`
}

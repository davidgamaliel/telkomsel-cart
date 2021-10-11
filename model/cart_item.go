package model

import "github.com/jinzhu/gorm"

type CartItem struct {
	gorm.Model
	ProductId     uint    `json:"product_id"`
	Quantity      int     `json:"quantity"`
	TransactionID uint    `json:"transaction_id"`
	UserId        uint    `json:"user_id"`
	Product       Product `json:"product"`
}

func (CartItem) TableName() string {
	return "cart_items"
}

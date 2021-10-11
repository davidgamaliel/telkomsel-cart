package model

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	State         string     `json:"state"`
	UserId        uint       `json:"user_id"`
	PaymentMethod string     `json:"payment_method"`
	TotalPrice    uint       `json:"total_price"`
	RemitAmount   uint       `json:"remit_amount"`
	CartItems     []CartItem `json:"cart_items"`
}

type TransactionCreateRequest struct {
	Transaction
	CartItemIds []uint `json:"cart_item_ids"`
}

func (Transaction) TableName() string {
	return "transactions"
}

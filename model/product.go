package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint64 `json:"price"`
}

func (Product) TableName() string {
	return "products"
}

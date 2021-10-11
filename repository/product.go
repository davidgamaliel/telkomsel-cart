package repository

import (
	"go-cart-api/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func (db *productRepository) GetProducts() (products []model.Product, err error) {
	return products, db.connection.Find(&products).Error
}

package repository

import (
	"go-cart-api/model"

	"github.com/jinzhu/gorm"
)

type CartItemRepository interface {
	CreateItem(model.CartItem) (model.CartItem, error)
	GetItemsByUserId(uint) ([]model.CartItem, error)
	GetItemsByIds([]uint) ([]model.CartItem, error)
	UpdateItem(model.CartItem) (model.CartItem, error)
}

type cartItemRepository struct {
	connection *gorm.DB
}

func NewCartItemRepository() CartItemRepository {
	return &cartItemRepository{
		connection: DB(),
	}
}

func (db *cartItemRepository) CreateItem(cartItem model.CartItem) (model.CartItem, error) {
	return cartItem, db.connection.Create(&cartItem).Error
}

func (db *cartItemRepository) GetItemsByUserId(userId uint) (cartItems []model.CartItem, err error) {
	return cartItems, db.connection.Preload("Product").Find(&cartItems, "user_id=?", userId).Error
}

func (db *cartItemRepository) GetItemsByIds(ids []uint) (cartItems []model.CartItem, err error) {
	return cartItems, db.connection.Preload("Product").Where(ids).Find(&cartItems).Error
}

func (db *cartItemRepository) UpdateItem(cartItem model.CartItem) (model.CartItem, error) {
	if err := db.connection.First(&cartItem, cartItem.ID).Error; err != nil {
		return cartItem, err
	}
	return cartItem, db.connection.Model(&cartItem).Updates(&cartItem).Error
}

package repository

import (
	"fmt"
	"go-cart-api/model"

	"github.com/jinzhu/gorm"
)

type TransactionRepository interface {
	CreateTransaction(model.Transaction) (model.Transaction, error)
	GetTransactionsByUserId(uint) ([]model.Transaction, error)
	GetTransactionsById(uint) (model.Transaction, error)
	UpdateTransaction(model.Transaction) (model.Transaction, error)
}

type transactionRepository struct {
	connection *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		connection: DB(),
	}
}

func (db *transactionRepository) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	return transaction, db.connection.Create(&transaction).Error
}

func (db *transactionRepository) GetTransactionsByUserId(userId uint) (transactions []model.Transaction, err error) {
	return transactions, db.connection.Preload("CartItems.Product").Find(&transactions, "user_id=?", userId).Error
}

func (db *transactionRepository) GetTransactionsById(id uint) (transaction model.Transaction, err error) {
	return transaction, db.connection.Preload("CartItems.Product").First(&transaction, id).Error
}

func (db *transactionRepository) UpdateTransaction(transaction model.Transaction) (model.Transaction, error) {
	fmt.Println(transaction.State)
	return transaction, db.connection.Model(&transaction).Updates(&transaction).Error
}

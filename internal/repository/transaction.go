package repository

import (
	"github.com/hameedhub/pismo/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction model.Transaction) (model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func (t transactionRepository) Create(transaction model.Transaction) (model.Transaction, error) {
	err := t.db.Create(&transaction).Error
	return transaction, err
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{db: db}

}

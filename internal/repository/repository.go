package repository

import "gorm.io/gorm"

type IRepository struct {
	AccountRepository     AccountRepository
	TransactionRepository TransactionRepository
}

func Init(db *gorm.DB) *IRepository {
	AccountRepository := NewAccountRepository(db)
	TransactionRepository := NewTransactionRepository(db)
	return &IRepository{AccountRepository: AccountRepository,
		TransactionRepository: TransactionRepository}
}

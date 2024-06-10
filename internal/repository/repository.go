package repository

import "gorm.io/gorm"

type IRepository struct {
	AccountRepository       AccountRepository
	TransactionRepository   TransactionRepository
	OperationTypeRepository OperationTypeRepository
}

func Init(db *gorm.DB) *IRepository {
	AccountRepository := NewAccountRepository(db)
	TransactionRepository := NewTransactionRepository(db)
	OperationTypeRepository := NewOperationTypeRepository(db)
	return &IRepository{AccountRepository: AccountRepository,
		TransactionRepository:   TransactionRepository,
		OperationTypeRepository: OperationTypeRepository}
}

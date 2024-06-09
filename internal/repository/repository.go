package repository

import "gorm.io/gorm"

type IRepository struct {
	AccountRepository AccountRepository
}

func Init(db *gorm.DB) *IRepository {
	AccountRepository := NewAccountRepository(db)
	return &IRepository{AccountRepository: AccountRepository}
}

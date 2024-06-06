package repository

import "gorm.io/gorm"

type IRepository struct {
	AccountRepo *AccountRepo
}

func Init(db *gorm.DB) *IRepository {
	accountRepo := NewAccountRepository(db)
	return &IRepository{AccountRepo: accountRepo}
}

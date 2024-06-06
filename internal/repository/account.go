package repository

import (
	"github.com/hameedhub/pismo/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetByID(id uint64) (model.Account, error)
	Create(user model.Account) (model.Account, error)
}

type AccountRepo struct {
	db *gorm.DB
}

func (a AccountRepo) GetByID(id uint64) (model.Account, error) {
	var account model.Account
	err := a.db.First(&account, id).Error
	return account, err
}

func (a AccountRepo) Create(account model.Account) (model.Account, error) {
	err := a.db.Create(&account).Error
	return account, err
}

func NewAccountRepository(db *gorm.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

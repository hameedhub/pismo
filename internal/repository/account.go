package repository

import (
	"github.com/hameedhub/pismo/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetByID(id int) (model.Account, error)
	GetByDocumentNumber(docNum string) (model.Account, error)
	Create(user model.Account) (model.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func (a accountRepository) GetByID(id int) (model.Account, error) {
	var account model.Account
	err := a.db.First(&account, id).Error
	return account, err
}

func (a accountRepository) GetByDocumentNumber(docNum string) (model.Account, error) {
	var account model.Account
	err := a.db.Where("document_number =?", docNum).First(&account).Error
	return account, err
}

func (a accountRepository) Create(account model.Account) (model.Account, error) {
	err := a.db.Create(&account).Error
	return account, err
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{
		db: db,
	}
}

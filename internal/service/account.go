package service

import (
	"errors"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/repository"
	"gorm.io/gorm"
)

type AccountService interface {
	Create(account model.Account) (model.Account, error)
	GetById(id uint64) (model.Account, error)
}

type accountService struct {
	repo *repository.IRepository
}

func (as accountService) Create(account model.Account) (model.Account, error) {
	existingAccount, err := as.repo.AccountRepository.GetByDocumentNumber(account.DocumentNumber)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return as.repo.AccountRepository.Create(account)
	}
	if err != nil {
		return existingAccount, err
	}
	return account, errors.New("document number already exist")
}

func (as accountService) GetById(id uint64) (model.Account, error) {
	return as.repo.AccountRepository.GetByID(id)
}

func NewAccountService(repo *repository.IRepository) AccountService {
	return &accountService{
		repo: repo,
	}
}

package service

import (
	"errors"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/repository"
	"gorm.io/gorm"
)

type TransactionService interface {
	Create(transaction model.Transaction) (model.Transaction, error)
}

type transactionService struct {
	repo *repository.IRepository
}

func (t transactionService) Create(transaction model.Transaction) (model.Transaction, error) {
	_, err := t.repo.AccountRepository.GetByID(transaction.AccountId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return transaction, err
	}
	//TODO:: Implement balance update
	return t.repo.TransactionRepository.Create(transaction)
}

func NewTransactionService(repo *repository.IRepository) TransactionService {
	return transactionService{repo: repo}
}

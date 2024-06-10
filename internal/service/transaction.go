package service

import (
	"errors"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/repository"
	"gorm.io/gorm"
	"math"
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
		return transaction, errors.New("account id not found")
	}
	opt, err := t.repo.OperationTypeRepository.GetById(transaction.OperationTypeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return transaction, errors.New("operation type id not found")
	}
	switch opt.ID {
	case model.PAYMENT:
		transaction.Amount = float32(math.Abs(float64(transaction.Amount)))
	case model.PURCHASE, model.INSTALLMENT_PURCHASE, model.WITHDRAWAL:
		transaction.Amount = (transaction.Amount) * -1
	default:
		return transaction, errors.New("operation type not define in constant")
	}

	//TODO:: Implement balance update
	return t.repo.TransactionRepository.Create(transaction)
}

func NewTransactionService(repo *repository.IRepository) TransactionService {
	return transactionService{repo: repo}
}

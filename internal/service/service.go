package service

import (
	"github.com/hameedhub/pismo/internal/repository"
)

type Service struct {
	AccountService     AccountService
	TransactionService TransactionService
}

func Init(repo *repository.IRepository) *Service {
	as := NewAccountService(repo)
	ts := NewTransactionService(repo)
	return &Service{AccountService: as,
		TransactionService: ts}
}

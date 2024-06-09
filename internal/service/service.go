package service

import (
	"github.com/hameedhub/pismo/internal/repository"
)

type Service struct {
	AccountService AccountService
}

func Init(repo *repository.IRepository) *Service {
	as := NewAccountService(repo)
	return &Service{AccountService: as}
}

package domain

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

var (
	ErrDocumentNumberEmpty = errors.New("document number can't be empty")
)

type Account struct {
	gorm.Model
	AccountID      uint64  `json:"account_id" gorm:"unique;primaryKey;autoIncrement"`
	DocumentNumber string  `json:"document_number"`
	Balance        float64 `json:"balance"`
}

func (a *Account) IsValid() error {
	if len(strings.TrimSpace(a.DocumentNumber)) == 0 {
		return ErrDocumentNumberEmpty
	}
	return nil
}

func (a *Account) NewAccount() Account {
	return Account{
		AccountID:      a.AccountID,
		DocumentNumber: a.DocumentNumber,
		Balance:        a.Balance,
	}
}

package model

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

var (
	ErrDocumentNumberInvalid = errors.New("invalid document number")
)

type Account struct {
	gorm.Model
	ID             int     `json:"account_id" gorm:"unique;primaryKey;autoIncrement"`
	DocumentNumber string  `json:"document_number"`
	Balance        float64 `json:"balance"`
}

func (a *Account) IsValid() error {
	_, err := strconv.Atoi(a.DocumentNumber)
	if len(a.DocumentNumber) < 3 || err != nil {
		return ErrDocumentNumberInvalid
	}
	return nil
}

func (a *Account) NewAccount() Account {
	return Account{
		ID:             a.ID,
		DocumentNumber: a.DocumentNumber,
		Balance:        a.Balance,
	}
}

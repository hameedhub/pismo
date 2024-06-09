package model

import (
	"gorm.io/gorm"
	"time"
)

const (
	PURCHASE             = 1
	INSTALLMENT_PURCHASE = 2
	WITHDRAWAL           = 3
	PAYMENT              = 4
)

type Transaction struct {
	gorm.Model
	ID              int       `json:"transaction_id" gorm:"unique;primaryKey;autoIncrement"`
	AccountId       int       `json:"account_id"`
	OperationTypeId int       `json:"operation_type_id"`
	CreatedAt       time.Time `json:"event_date"`
}

func (a *Transaction) NewTransaction() Transaction {
	return Transaction{
		ID:              a.ID,
		AccountId:       a.AccountId,
		OperationTypeId: a.OperationTypeId,
	}
}

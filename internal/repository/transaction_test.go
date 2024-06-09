package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestAccountRepository_Create(t *testing.T) {
	db, mock, gormDb := SetupMock(t)
	defer db.Close()

	newTransaction := model.Transaction{
		OperationTypeId: 1,
		AccountId:       1,
	}
	expectedTransaction := newTransaction
	expectedTransaction.ID = 1

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "transactions" ("created_at","updated_at","deleted_at","account_id","operation_type_id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			1,
			1,
		).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedTransaction.ID))
	mock.ExpectCommit()
	repo := NewTransactionRepository(gormDb)
	transaction, err := repo.Create(newTransaction)
	assert.Nil(t, err)
	assert.Equal(t, expectedTransaction.ID, transaction.ID)

}

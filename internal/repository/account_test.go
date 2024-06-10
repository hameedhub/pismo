package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error")
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		t.Errorf("error opening gorm db: %s", err)
	}
	return db, mock, gormDb
}

func TestAccountRepo_GetByID(t *testing.T) {
	db, mock, gormDb := SetupMock(t)
	defer db.Close()

	id := 1
	expected := model.Account{ID: id, DocumentNumber: "3382098032", Balance: 0}
	rows := sqlmock.NewRows([]string{"id", "document_number", "balance"}).AddRow(expected.ID, expected.DocumentNumber, expected.Balance)
	mock.ExpectQuery("^SELECT \\* FROM \"accounts\" WHERE \"accounts\"\\.\"id\" = \\$1 AND \"accounts\"\\.\"deleted_at\" IS NULL ORDER BY \"accounts\"\\.\"id\" LIMIT \\$2").WithArgs(id, 1).WillReturnRows(rows)
	repo := NewAccountRepository(gormDb)
	account, err := repo.GetByID(id)
	assert.Nil(t, err)
	assert.Equal(t, expected, account)
}

func TestAccountRepo_Create(t *testing.T) {
	db, mock, gormDb := SetupMock(t)
	defer db.Close()

	newAccount := model.Account{
		DocumentNumber: "3382098032",
		Balance:        0.0,
	}
	expectedAccount := newAccount
	expectedAccount.ID = 1 // Assume the database assigns ID 1

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts" ("created_at","updated_at","deleted_at","document_number","balance") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			"3382098032",
			0.0,
		).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedAccount.ID))
	mock.ExpectCommit()

	// Test Create
	repo := NewAccountRepository(gormDb)
	account, err := repo.Create(newAccount)

	assert.Nil(t, err)
	assert.Equal(t, expectedAccount.ID, account.ID)

}

package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestInit(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal("db mock error")
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	repo := Init(gormDb)

	if repo == nil || repo.AccountRepository == nil {
		t.Fatalf("fail to initialize repo.")
	}
}

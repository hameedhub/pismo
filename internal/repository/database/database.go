package database

import (
	"fmt"
	"github.com/hameedhub/pismo/internal/model"
	"github.com/hameedhub/pismo/internal/repository/database/seeder"

	"github.com/hameedhub/pismo/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Uri(c config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func Run(c config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(Uri(c)), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	err = db.AutoMigrate(&model.Account{}, &model.Transaction{}, &model.OperationType{})
	if err != nil {
		log.Fatal("Migration error: ", err)
	}
	err = seeder.OperationTypes(db)
	if err != nil {
		log.Fatal("Operation types seeds error: ", err)
	}

	return db
}

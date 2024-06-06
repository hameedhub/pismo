package database

import (
	"fmt"
	"github.com/hameedhub/pismo/internal/account/domain"
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
	err = db.AutoMigrate(&domain.Account{})
	if err != nil {
		log.Fatal("Migration error: ", err)
	}

	return db
}

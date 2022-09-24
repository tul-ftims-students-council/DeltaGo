package db

import (
	"log"

	"delta-go/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	url := "postgres://postgres:d9856ae29ffe010d93230ff6b9f310541c699934893e0e53@delta-go-db.internal:8080/delta"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}

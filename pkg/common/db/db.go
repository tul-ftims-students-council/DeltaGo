package db

import (
	"log"

	"delta-go/pkg/common/config"
	"delta-go/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.DBurl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}

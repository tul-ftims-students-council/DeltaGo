package db

import (
	"log"

	"delta-go/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	url := "postgres://user:n9GQTJC95CmxKE9D4CMpvXqeFPLKaz5Y@dpg-ccnh1qta49940mojcc0g-a.oregon-postgres.render.com/delta_6y61"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}

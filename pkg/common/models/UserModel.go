package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

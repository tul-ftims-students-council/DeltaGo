package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Place struct {
	gorm.Model
	UserEmail      string    `json:"user_email"`
	DateTillExpire time.Time `json:"till_expire"`
	IsSold         bool      `json:"is_sold"`
}

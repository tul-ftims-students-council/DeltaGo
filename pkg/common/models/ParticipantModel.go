package models

import (
	"github.com/jinzhu/gorm"
)

type Participant struct {
	gorm.Model

	Id             int    `json:"id" gorm:"primaryKey"`
	UserEmail      string `json:"user_email"`
	Major          string `json:"major"`
	Year           int    `json:"year"`
	TShirtSize     string `json:"t_shirt_size"`
	Diet           string `json:"diet"`
	PaymentFile    []byte `json:"payment_file" gorm:"type:bytea"`
	FileExtension  string `json:"file_extension"`
	InvoiceAddress string `json:"invoice_address"`
	FootSize       string `json:"footSize"`
}

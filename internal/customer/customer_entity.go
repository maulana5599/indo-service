package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id           int            `json:"id" gorm:"primary_key"`
	UserId       int            `json:"user_id"`
	Alamat       string         `json:"alamat"`
	NoTelp       string         `json:"no_telp"`
	TempatLahir  string         `json:"tempat_lahir"`
	TanggalLahir time.Time      `json:"tanggal_lahir"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type CustomerV struct {
	Id           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
	Alamat       string `json:"alamat"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
}

func (Customer) TableName() string {
	return "customer_m"
}

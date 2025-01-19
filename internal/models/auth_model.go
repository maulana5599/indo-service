package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetUsername(email string) (entity.EmailDatabase, error) {
	var result entity.EmailDatabase
	tx := config.DB.Table("users").Where("email = ?", email).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

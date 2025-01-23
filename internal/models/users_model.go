package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetUsers() ([]entity.Users, error) {
	var result []entity.Users
	tx := config.DB.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetUsersById(id int) (entity.Users, error) {
	var result entity.Users
	tx := config.DB.Where("id = ?", id).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

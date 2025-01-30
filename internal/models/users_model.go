package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
	"strings"
)

func GetUsers() ([]entity.Users, error) {
	var result []entity.Users
	tx := config.DB.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetCustomerAll() ([]entity.Customer, error) {
	var result []entity.Customer
	tx := config.DB.Table("customer_m").Where("deleted_at IS NULL").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func SearchCustomer(name string) ([]entity.Customer, error) {
	var result []entity.Customer
	tx := config.DB.Table("customer_m").Where("deleted_at IS NULL").Where("LOWER(nama) LIKE ?", "%"+strings.ToLower(name)+"%").Find(&result)
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

package auth

import (
	"echo-boilerplate/config"
	customerEntity "echo-boilerplate/internal/customer"

	"gorm.io/gorm"
)

func CreateNewUser(user Users, customer customerEntity.Customer) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		customer.UserId = user.Id
		if err := tx.Create(&customer).Error; err != nil {
			return err
		}

		roleCustomer := Role{
			UserId: user.Id,
			RoleId: 1,
		}

		if err := tx.Table("role_mp").Create(&roleCustomer).Error; err != nil {
			return err
		}

		return nil
	})
}

func CheckEmail(email string) (EmailDatabase, error) {
	var result EmailDatabase
	tx := config.DB.Table("users").Where("email = ?", email).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

func GetUsername(email string) (EmailDatabase, error) {
	var result EmailDatabase
	tx := config.DB.Table("users").Where("email = ?", email).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

func GetRole(userId int) ([]Role, error) {
	var result []Role
	tx := config.DB.Table("role_mp").Where("user_id = ?", userId).Find(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

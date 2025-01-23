package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetPembayaranCicilanId(userId int) ([]entity.PembayaranCicilan, error) {
	var result []entity.PembayaranCicilan
	tx := config.DB.Where("user_id = ?", userId).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

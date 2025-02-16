package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetMapel() ([]entity.Mapel, error) {
	var result []entity.Mapel
	tx := config.DB.Where("deleted_at IS NULL").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

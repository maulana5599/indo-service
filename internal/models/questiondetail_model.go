package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetDetailQuestion(questionDetailId int) (entity.QustionDetailM, error) {
	var result entity.QustionDetailM
	tx := config.DB.Where("questiondetail_id = ? AND deleted_at IS NULL", questionDetailId).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

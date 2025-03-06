package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetDetailQuiz(quizDetailId int) (entity.QuizDetailT, error) {
	var result entity.QuizDetailT
	tx := config.DB.Where("id = ? AND deleted_at IS NULL", quizDetailId).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

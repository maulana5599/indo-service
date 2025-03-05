package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetLearningTopic(roomId int) ([]entity.TopiclessonT, error) {
	var result []entity.TopiclessonT
	tx := config.DB.Where("room_id = ? AND deleted_at IS NULL", roomId).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetRoomTopic(roomId int) (entity.RoomTopic, error) {
	var result entity.RoomTopic
	tx := config.DB.Table("room_m").
		Select("room_m.*, major_m.major_name, users.name").
		Joins("JOIN users on users.id = room_m.teacher_id").
		Joins("JOIN major_m on major_m.major_id = room_m.major_id").
		Where("room_id = ?", roomId).First(&result)

	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

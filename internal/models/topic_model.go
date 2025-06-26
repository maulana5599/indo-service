package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetLearningTopic(roomId int, page int, perPage int, search string) ([]entity.TopiclessonT, int64, error) {
	var result []entity.TopiclessonT
	var totalRecord int64
	tx := config.DB.Where("room_id = ? AND deleted_at IS NULL AND is_archive IS FALSE", roomId).
		Offset((page - 1) * perPage).
		Limit(perPage)

	if search != "" {
		tx = tx.Where("topic_name ILIKE ?", "%"+search+"%")
	}

	tx.Find(&result)

	config.DB.Table("topiclesson_t").Where("room_id = ? AND deleted_at IS NULL AND is_archive IS FALSE", roomId).Count(&totalRecord)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	return result, totalRecord, nil
}

func GetRoomTopic(roomId int) (entity.RoomTopic, error) {
	var result entity.RoomTopic
	tx := config.DB.Table("room_m").
		Select("room_m.*, major_m.major_name, users.name").
		Joins("JOIN users on users.id = room_m.teacher_id").
		Joins("JOIN major_m on major_m.major_id = room_m.major_id").
		Where("room_id = ? AND room_m.deleted_at IS NULL", roomId).First(&result)

	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

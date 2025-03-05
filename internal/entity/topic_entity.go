package entity

import "time"

type TopiclessonT struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	TopicName   string    `json:"topic_name"`
	TopicDesc   string    `json:"topic_desc"`
	RoomId      int       `json:"room_id"`
	TeacherId   int       `json:"teacher_id"`
	IsCopy      bool      `json:"is_copy"`
	LastTopicId int       `json:"last_topic_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RoomTopic struct {
	RoomId    int    `json:"room_id" gorm:"primaryKey"`
	RoomName  string `json:"room_name"`
	MajorName string `json:"major_name"`
	MajorId   int    `json:"major_id"`
	Name      string `json:"name"`
}

func (topic *TopiclessonT) TableName() string {
	return "topiclesson_t"
}

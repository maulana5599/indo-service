package entity

import "time"

type Mapel struct {
	MajorId   int       `json:"major_id" gorm:"primaryKey"`
	MajorName string    `json:"mata_pelajaran" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Mapel) TableName() string {
	return "major_m"
}

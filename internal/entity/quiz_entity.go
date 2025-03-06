package entity

type QuizDetailT struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	QuizId           int    `json:"quiz_id"`
	QuestiondetailId int    `json:"questiondetail_id"`
	Answer           string `json:"answer"`
	UserId           int    `json:"user_id"`
	IsSelesai        bool   `json:"is_selesai"`
}

func (QuizDetailT) TableName() string {
	return "quizdetail_t"
}

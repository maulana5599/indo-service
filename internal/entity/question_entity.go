package entity

type QustionDetailM struct {
	QuestiondetailId int    `json:"questiondetail_id" gorm:"primaryKey"`
	QuestionId       int    `json:"question_id"`
	Question         string `json:"question"`
	Answer           string `json:"answer"`
	AnswerA          string `json:"answer_a"`
	AnswerB          string `json:"answer_b"`
	AnswerC          string `json:"answer_c"`
	AnswerD          string `json:"answer_d"`
	AnswerE          string `json:"answer_e"`
	Section          int    `json:"section"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func (QustionDetailM) TableName() string {
	return "questiondetail_m"
}

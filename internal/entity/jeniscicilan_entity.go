package entity

import "time"

type JenisCicilan struct {
	JenispinjamanId int `gorm:"primary_key"`
	NamaCicilan     string
	PokokCicilan    float64
	TotalAngsuran   float64
	JumlahAngsuran  int
	MarginCicilan   float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (JenisCicilan) TableName() string {
	return "jeniscicilan_m"
}

type JenisCicilanRequest struct {
	NamaCicilan    string  `json:"nama_cicilan"`
	PokokCicilan   float64 `json:"pokok_cicilan"`
	TotalAngsuran  float64 `json:"total_angsuran"`
	JumlahAngsuran int     `json:"jumlah_angsuran"`
	MarginCicilan  float64 `json:"margin_cicilan"`
}

type JenisCicilanResponse struct {
	JenispinjamanId int     `json:"jenispinjaman_id"`
	NamaCicilan     string  `json:"nama_cicilan"`
	PokokCicilan    float64 `json:"pokok_cicilan"`
	TotalAngsuran   float64 `json:"total_angsuran"`
	JumlahAngsuran  int     `json:"jumlah_angsuran"`
	MarginCicilan   float64 `json:"margin_cicilan"`
}

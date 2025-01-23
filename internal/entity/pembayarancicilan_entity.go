package entity

import "time"

type PembayaranCicilan struct {
	PembayarancicilanId int `gorm:"primary_key"`
	PengajuancicilanId  int
	JeniscicilanId      int
	Angsuran            int
	NominalPembayaran   float64
	StatusPembayaran    int
	UserId              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type PembayaranCicilanResponse struct {
	PembayarancicilanId int       `json:"pembayarancicilan_id"`
	PengajuancicilanId  int       `json:"pengajuancicilan_id"`
	JeniscicilanId      int       `json:"jeniscicilan_id"`
	Angsuran            int       `json:"angsuran"`
	NominalPembayaran   int       `json:"nominal_pembayaran"`
	StatusPembayaran    int       `json:"status_pembayaran"`
	UserId              int       `json:"user_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type StatusPembayaranRequest struct {
	StatusPembayaran    int `json:"status_pembayaran"`
	PembayarancicilanId int `json:"pembayarancicilan_id"`
}

func (PembayaranCicilan) TableName() string {
	return "pembayarancicilan_t"
}

package entity

import (
	"time"
)

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

type PembayaranCicilanView struct {
	PembayarancicilanId int        `json:"pembayarancicilan_id"`
	PengajuancicilanId  int        `json:"pengajuancicilan_id"`
	Angsuran            int        `json:"angsuran"`
	NominalPembayaran   float64    `json:"nominal_pembayaran"`
	StatusPembayaran    int        `json:"status_pembayaran"`
	NamaCicilan         string     `json:"nama_cicilan"`
	TotalAngsuran       float64    `json:"total_angsuran"`
	MarginCicilan       float64    `json:"margin_cicilan"`
	UserId              int        `json:"user_id"`
	Nama                string     `json:"nama_user"`
	PaymentAt           *time.Time `json:"payment_at"`
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
	StatusPembayaran    int    `json:"status_pembayaran"`
	NominalPembayaran   int    `json:"nominal_pembayaran"`
	Keterangan          string `json:"keterangan"`
	PembayarancicilanId int    `json:"pembayarancicilan_id"`
}

func (PembayaranCicilan) TableName() string {
	return "pembayarancicilan_t"
}

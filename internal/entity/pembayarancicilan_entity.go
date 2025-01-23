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

func (PembayaranCicilan) TableName() string {
	return "pembayarancicilan_t"
}

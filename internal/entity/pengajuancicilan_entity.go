package entity

import "time"

type PengajuanCicilan struct {
	PengajuancicilanId int `gorm:"primary_key"`
	UserId             int
	NoKtp              string
	Alamat             string
	NoSiswa            string
	Pekerjaan          string
	Orangtua           string
	NohpOrtu           string
	KontakDarurat      string
	Jaminan            string
	TglPencairan       time.Time `gorm:"default:null"`
	Keterangan         string    `gorm:"default:null"`
	JeniscicilanId     int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (PengajuanCicilan) TableName() string {
	return "pengajuancicilan_t"
}

type PengajuanCicilanView struct {
	Nama               string     `json:"nama_user"`
	PengajuancicilanId int        `json:"pengajuancicilan_id"`
	UserId             int        `json:"user_id"`
	NoKtp              string     `json:"no_ktp"`
	Alamat             string     `json:"alamat"`
	Pekerjaan          string     `json:"pekerjaan"`
	Orangtua           string     `json:"orang_tua"`
	NohpOrtu           string     `json:"nohp_ortu"`
	KontakDarurat      string     `json:"kontak_darurat"`
	Jaminan            string     `json:"jaminan"`
	TglPencairan       *time.Time `json:"tgl_pencairan"`
	JenisCicilanid     int        `json:"jeniscicilan_id"`
	NamaCicilan        string     `json:"nama_cicilan"`
	PokokCicilan       float64    `json:"pokok_cicilan"`
	TotalAngsuran      float64    `json:"total_angsuran"`
	JumlahAngsuran     int        `json:"jumlah_angsuran"`
}

type PengajuanCicilanRequest struct {
	UserId        int    `json:"user_id"`
	NoKtp         string `json:"no_ktp"`
	Alamat        string `json:"alamat"`
	NoSiswa       string `json:"no_siswa"`
	Pekerjaan     string `json:"pekerjaan"`
	OrangTua      string `json:"orang_tua"`
	NohpOrtu      string `json:"nohp_ortu"`
	KontakDarurat string `json:"kontak_darurat"`
	Jaminan       string `json:"jaminan"`
	JenisCicilan  int    `json:"jenis_cicilan"`
}

type PengajuanCicilanResponse struct {
	PengajuancicilanId int         `json:"pengajuancicilan_id"`
	UserId             interface{} `json:"user_id,omitempty"`
	NoKtp              string      `json:"no_ktp"`
	Alamat             string      `json:"alamat"`
	NoSiswa            string      `json:"no_siswa"`
	Pekerjaan          string      `json:"pekerjaan"`
	Orangtua           string      `json:"orang_tua"`
	NohpOrtu           string      `json:"nohp_ortu"`
	KontakDarurat      string      `json:"kontak_darurat"`
	Jaminan            string      `json:"jaminan"`
	TglPencairan       *string     `json:"tgl_pencairan"`
	Keterangan         *string     `json:"keterangan"`
	JeniscicilanId     int         `json:"jeniscicilan_id"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`
}

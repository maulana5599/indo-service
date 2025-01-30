package entity

import "time"

type Users struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Users) TableName() string {
	return "users"
}

type ResponseUsers struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Customer struct {
	Id             int       `json:"id" gorm:"primary_key"`
	Nama           string    `json:"nama_user"`
	Alamat         string    `json:"alamat"`
	NamaOrangtua   string    `json:"nama_orangtua"`
	NoTelp         string    `json:"no_telp"`
	NoTelpOrangtua string    `json:"no_telp_orangtua"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tanggal_lahir"`
	JenisKelamin   string    `json:"jenis_kelamin"`
	CreatedAt      time.Time `json:"created_at"`
}

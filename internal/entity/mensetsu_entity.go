package entity

import "time"

type MensetsuV struct {
	Hari              string
	TanggalMensetsu   time.Time
	NamaPerusahaan    int
	Penempatan        string
	JumlahKandidat    int
	Link              string
	Status            int
	JobId             int
	KumiaiId          int
	TanggalPengumuman time.Time
	JkKandidat        int
	IsOnline          bool
	Gaji              float64
	Cv                int
	PerusahaanNama    string
	JobName           string
	KumiaiName        string
	JumlahSiswa       int
	CheckCv           bool
}

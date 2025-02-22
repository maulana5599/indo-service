package entity

type GrafikJobs struct {
	JenisPekerjaan string `json:"jenis_pekerjaan"`
	TotalPekerjaan int    `json:"total_pekerjaan"`
}

type GrafikAngkatan struct {
	Angkatan      string `json:"angkatan"`
	TotalAngkatan int    `json:"total_angkatan"`
}

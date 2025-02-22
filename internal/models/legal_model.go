package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
)

func GetGrafikJobs() ([]entity.GrafikJobs, error) {
	var result []entity.GrafikJobs
	tx := config.DB.Raw(`SELECT 
			lm.lookup_name as jenis_pekerjaan,
			COUNT(pst.jenispekerjaan_id) as total_pekerjaan 
		FROM pekerjaan_siswa_t pst
		JOIN lookup_m lm on lm.lookup_id = pst.jenispekerjaan_id 
		JOIN users on users.id = pst.siswa_id
		WHERE pst.deleted_at is null
		GROUP BY pst.jenispekerjaan_id, lm.lookup_name`).First(&result)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetGrafikJobsAngkatan(perusahaanIdInt int) ([]entity.GrafikAngkatan, error) {
	var result []entity.GrafikAngkatan
	tx := config.DB.Raw(`SELECT 
		angkatan_kumiai_t.angkatan,
		COUNT(angkatan_id) as total_angkatan
	FROM perusahaankumiai_t
	JOIN angkatan_kumiai_t on angkatan_kumiai_t.id = perusahaankumiai_t.angkatan_id 
	WHERE perusahaankumiai_t.deleted_at is null
	and perusahaankumiai_t.kumiaidetail_id = ?
	GROUP BY perusahaankumiai_t.angkatan_id,angkatan_kumiai_t.angkatan, perusahaankumiai_t.kumiaidetail_id`, perusahaanIdInt).First(&result)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

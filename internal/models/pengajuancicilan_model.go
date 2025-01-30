package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
	"errors"
	"time"
)

func GetDataPengajuanCicilan() ([]entity.PengajuanCicilanView, error) {
	var result []entity.PengajuanCicilanView
	tx := config.DB.Table("pengajuan_v").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func AddCicilan(pengajuanCicilan *entity.PengajuanCicilanRequest) error {
	tx := config.DB.Begin()

	pengajuanData := entity.PengajuanCicilan{
		UserId:         pengajuanCicilan.UserId,
		NoKtp:          pengajuanCicilan.NoKtp,
		Alamat:         pengajuanCicilan.Alamat,
		NoSiswa:        pengajuanCicilan.NoSiswa,
		Pekerjaan:      pengajuanCicilan.Pekerjaan,
		Orangtua:       pengajuanCicilan.OrangTua,
		NohpOrtu:       pengajuanCicilan.NohpOrtu,
		KontakDarurat:  pengajuanCicilan.KontakDarurat,
		Jaminan:        pengajuanCicilan.Jaminan,
		JeniscicilanId: pengajuanCicilan.JenisCicilan,
	}

	createPengajuanCicilan := tx.Create(&pengajuanData)

	if createPengajuanCicilan.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	var JenisCicilan entity.JenisCicilan
	errFind := tx.Where("jenispinjaman_id = ?", pengajuanCicilan.JenisCicilan).First(&JenisCicilan)

	if errFind.Error != nil {
		tx.Rollback()
		return errFind.Error
	}

	var PembayaranCicilan []entity.PembayaranCicilan
	var JumlahAngsuran int = JenisCicilan.JumlahAngsuran

	for i := 1; i <= int(JumlahAngsuran); i++ {
		PembayaranCicilan = append(PembayaranCicilan, entity.PembayaranCicilan{
			PengajuancicilanId: pengajuanData.PengajuancicilanId,
			JeniscicilanId:     pengajuanCicilan.JenisCicilan,
			Angsuran:           i,
			NominalPembayaran:  JenisCicilan.PokokCicilan,
			StatusPembayaran:   0,
			UserId:             pengajuanCicilan.UserId,
		})
	}

	tx.Create(&PembayaranCicilan)

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	tx.Commit()
	return nil
}

func BatalCicilan(pengajuanId int, keterangan string) error {
	tx := config.DB.Begin()

	notExistData := tx.Table("pengajuancicilan_t").
		Where("pengajuancicilan_id = ?", pengajuanId).
		Where("deleted_at IS NULL").
		First(&entity.PengajuanCicilan{})

	if notExistData.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("data pengajuan cicilan tidak ditemukan")
	}

	hapusHeader := tx.Table("pengajuancicilan_t").
		Where("pengajuancicilan_id = ?", pengajuanId).
		Where("deleted_at IS NULL").
		Updates(map[string]interface{}{
			"keterangan": keterangan,
			"deleted_at": time.Now().Format("2006-01-02 15:04:05"),
		})

	if hapusHeader.Error != nil {
		tx.Rollback()
		return hapusHeader.Error
	}

	hapusDetail := tx.Table("pembayarancicilan_t").
		Where("pengajuancicilan_id = ? AND deleted_at IS NULL", pengajuanId).
		Update("deleted_at", time.Now())

	if hapusDetail.Error != nil {
		tx.Rollback()
		return hapusDetail.Error
	}

	tx.Commit()
	return nil
}

func GetCicilanUser(id int) ([]entity.PengajuanCicilan, error) {
	var result []entity.PengajuanCicilan
	tx := config.DB.Where("user_id = ?", id).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

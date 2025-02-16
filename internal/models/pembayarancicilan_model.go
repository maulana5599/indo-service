package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
	"errors"
	"time"
)

func GetHeaderCicilanId(userId int) ([]entity.PengajuanCicilanView, error) {
	var result []entity.PengajuanCicilanView
	tx := config.DB.Table("pengajuan_v").Where("user_id = ?", userId).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetPembayaranCicilanId(userId int) ([]entity.PembayaranCicilanView, error) {
	var result []entity.PembayaranCicilanView
	tx := config.DB.Table("pembayarancicilan_v").Where("pengajuancicilan_id = ?", userId).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetPembayaranDetailId(pembayaranCicilanId int) (entity.PembayaranCicilanView, error) {
	var result entity.PembayaranCicilanView
	tx := config.DB.Table("pembayarancicilan_v").Where("pembayarancicilan_id = ?", pembayaranCicilanId).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

func UbahStatusPembayaranCicilan(pembayaranCicilan *entity.StatusPembayaranRequest) error {
	tx := config.DB.Begin()

	notExistData := tx.Table("pembayarancicilan_t").
		Where("pembayarancicilan_id = ?", pembayaranCicilan.PembayarancicilanId).
		Where("deleted_at IS NULL").
		First(&entity.PembayaranCicilan{})

	if notExistData.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("data cicilan tidak ditemukan")
	}

	updatePembayaranCicilan := tx.Table("pembayarancicilan_t").
		Where("pembayarancicilan_id = ?", pembayaranCicilan.PembayarancicilanId).
		Updates(map[string]interface{}{
			"status_pembayaran": pembayaranCicilan.StatusPembayaran,
			"keterangan":        pembayaranCicilan.Keterangan,
			"payment_at":        time.Now().Format("2006-01-02 15:04:05"),
		})

	if updatePembayaranCicilan.Error != nil {
		tx.Rollback()
		return updatePembayaranCicilan.Error
	}

	tx.Commit()
	return nil
}

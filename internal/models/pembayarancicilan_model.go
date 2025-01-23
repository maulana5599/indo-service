package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
	"errors"
)

func GetPembayaranCicilanId(userId int) ([]entity.PembayaranCicilanResponse, error) {
	var result []entity.PembayaranCicilanResponse
	tx := config.DB.Table("pembayarancicilan_t").Where("user_id = ?", userId).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
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
		})

	if updatePembayaranCicilan.Error != nil {
		tx.Rollback()
		return updatePembayaranCicilan.Error
	}

	tx.Commit()
	return nil
}

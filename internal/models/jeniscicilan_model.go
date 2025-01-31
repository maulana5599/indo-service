package models

import (
	"echo-boilerplate/config"
	"echo-boilerplate/internal/entity"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func GetJenisCicilan() ([]entity.JenisCicilan, error) {
	var result []entity.JenisCicilan
	tx := config.DB.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func GetJenisCicilanId(id int) (entity.JenisCicilan, error) {
	var result entity.JenisCicilan
	tx := config.DB.Where("jenispinjaman_id = ?", id).First(&result)
	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

func AddJenisCicilan(jenisCicilan *entity.JenisCicilan) error {
	tx := config.DB.Create(jenisCicilan)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func DeleteCicilan(id int) error {
	tx := config.DB.Delete(&entity.JenisCicilan{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func isNameUniqueDB(name string) bool {
	var count int64
	config.DB.Model(entity.JenisCicilan{}).Where("nama_cicilan = ?", name).Count(&count)
	return count == 0
}

func UniqueNameDB() validation.Rule {
	return validation.By(func(value interface{}) error {
		nama_cicilan, _ := value.(string)
		if !isNameUniqueDB(nama_cicilan) {
			return errors.New("nama cicilan sudah digunakan")
		}
		return nil
	})
}

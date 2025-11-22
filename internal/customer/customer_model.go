package customer

import (
	"echo-boilerplate/config"
	"strings"
)

func FindAllCustomers() ([]CustomerV, error) {
	var result []CustomerV
	tx := config.DB.Table("customer_v").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func FindCustomerByID(id int) (CustomerV, error) {
	var result CustomerV
	tx := config.DB.Table("customer_v").Where("id = ?", id).First(&result)
	if tx.Error != nil {
		return CustomerV{}, tx.Error
	}

	return result, nil
}

func SearchCustomer(name string) ([]Customer, error) {
	var result []Customer
	tx := config.DB.Table("customer_m").Where("deleted_at IS NULL").Where("LOWER(nama) LIKE ?", "%"+strings.ToLower(name)+"%").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func DeleteCustomer(id int) error {
	return config.DB.Delete(&Customer{}, "user_id = ?", id).Error
}

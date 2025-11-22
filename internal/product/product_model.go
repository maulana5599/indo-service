package product

import "echo-boilerplate/config"

func FindAllProducts(searchParam string) []Product {
	var result []Product

	tx := config.DB.Where("deleted_at IS NULL")

	if searchParam != "" {
		tx = tx.Where("product_name ILIKE ?", "%"+searchParam+"%")
	}

	tx.Find(&result)

	if tx.Error != nil {
		return nil
	}

	return result
}

func FindProductByID(id string) (Product, error) {
	var result Product
	tx := config.DB.Where("id = ?", id).First(&result)
	if tx.Error != nil {
		return Product{}, tx.Error
	}

	return result, nil
}

func FindProductByName(name string) (Product, error) {
	var result Product
	tx := config.DB.Where("product_name ILIKE ?", "%"+name+"%").First(&result)
	if tx.Error != nil {
		return Product{}, tx.Error
	}

	return result, nil
}

func CreateProduct(product Product) error {
	return config.DB.Create(&product).Error
}

func DeleteProductById(id string) error {
	return config.DB.Where("id = ?", id).Delete(&Product{}).Error
}

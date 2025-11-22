package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID              string         `json:"id"`
	ProductName     string         `json:"product_name"`
	ProductPrice    float32        `json:"product_price"`
	Brand           string         `json:"brand"`
	ProductInfo     string         `json:"product_info"`
	ProductImageUrl string         `json:"product_image_url"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func (Product) TableName() string {
	return "product_m"
}

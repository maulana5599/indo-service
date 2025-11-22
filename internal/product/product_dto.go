package product

type ProductRequest struct {
	ProductName     string  `json:"product_name"`
	ProductPrice    float32 `json:"product_price"`
	Brand           string  `json:"brand"`
	ProductInfo     string  `json:"product_info"`
	ProductImageUrl string  `json:"product_image_url"`
}

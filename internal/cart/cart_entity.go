package cart

type Cart struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (Cart) TableName() string {
	return "bucket_t"
}

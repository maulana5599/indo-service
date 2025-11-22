package cart

type CartRequest struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

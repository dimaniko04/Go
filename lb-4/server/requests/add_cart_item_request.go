package requests

type AddCartItemRequest struct {
	Quantity  int `json:"quantity"`
	ProductId int `json:"product_id"`
}

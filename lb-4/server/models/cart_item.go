package models

type CartItem struct {
	Id        int     `json:"id"`
	UserId    int     `json:"user_id"`
	Quantity  int     `json:"quantity"`
	ProductId int     `json:"product_id"`
	Product   Product `json:"product"`
}

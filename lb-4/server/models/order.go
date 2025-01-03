package models

type Order struct {
	Id        int         `json:"id"`
	UserId    int         `json:"user_id"`
	CreatedAt string      `json:"created_at"`
	Items     []OrderItem `json:"items"`
}

type OrderItem struct {
	Id        int     `json:"id"`
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product"`
}

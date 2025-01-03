package models

type Cart struct {
	Id        int        `json:"id"`
	UserId    int        `json:"user_id"`
	CartItems []CartItem `json:"-"`
}

type CartItem struct {
	Id        int     `json:"id"`
	Quantity  int     `json:"quantity"`
	ProductId int     `json:"user_id"`
	Product   Product `json:"-"`
}

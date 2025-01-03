package requests

type AddProductRequest struct {
	Name        string  `form:"name"`
	Description string  `form:"description"`
	Price       float64 `form:"price"`
	Stock       int     `form:"stock"`
}

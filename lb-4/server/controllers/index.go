package controllers

import "github.com/dimaniko04/Go/lb-4/server/services"

type Controllers struct {
	ProductController ProductController
}

func GetControllers(services *services.Services) *Controllers {
	return &Controllers{
		ProductController: &productController{productService: services.ProductService},
	}
}

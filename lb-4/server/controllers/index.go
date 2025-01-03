package controllers

import (
	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/dimaniko04/Go/lb-4/server/services"
)

type Controllers struct {
	ProductController ProductController
	AuthController    AuthController
	CartController    CartController
	OrderController   OrderController
}

func GetControllers(services *services.Services, env *config.Env) *Controllers {
	return &Controllers{
		ProductController: &productController{productService: services.ProductService},
		AuthController:    &authController{authService: services.AuthService},
		CartController: &cartController{
			cartService:     services.CartService,
			paymentService:  services.PaymentService,
			liqpayPublicKey: env.PublicKey,
		},
		OrderController: &orderController{orderService: services.OrderService},
	}
}

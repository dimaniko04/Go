package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/dimaniko04/Go/lb-4/server/middlewares"
	"github.com/gin-gonic/gin"
)

func orderRoutes(r *gin.RouterGroup, c controllers.OrderController, env *config.Env) {
	g := r.Group("/orders")

	g.Use(middlewares.AuthMiddleware([]string{"user"}, env.JwtSecret))
	{
		g.POST("/", c.GetUserOrders)
	}
}

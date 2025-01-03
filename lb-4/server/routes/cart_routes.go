package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/dimaniko04/Go/lb-4/server/middlewares"
	"github.com/gin-gonic/gin"
)

func cartRoutes(r *gin.RouterGroup, c controllers.CartController, env *config.Env) {
	g := r.Group("/cart")

	g.Use(middlewares.AuthMiddleware([]string{"user"}, env.JwtSecret))
	{
		g.GET("/", c.GetAll)
		g.POST("/", c.Add)
		g.DELETE("/:id", c.Delete)
		g.POST("/checkout", c.Checkout)
	}
}

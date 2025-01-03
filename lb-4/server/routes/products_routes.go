package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/dimaniko04/Go/lb-4/server/middlewares"
	"github.com/gin-gonic/gin"
)

func productRoutes(r *gin.RouterGroup, c controllers.ProductController, env *config.Env) {
	g := r.Group("/products")

	g.Use(middlewares.AuthMiddleware([]string{"admin"}, env.JwtSecret))
	{
		g.POST("/", c.Add)
		g.DELETE("/:id", c.Delete)
	}
	g.Use(middlewares.AuthMiddleware([]string{"admin", "user"}, env.JwtSecret))
	{
		g.GET("/", c.GetAll)
	}
}

package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.Engine, c *controllers.Controllers, env *config.Env) *gin.RouterGroup {
	router := api.Group("/api")

	authRoutes(router, c.AuthController)
	productRoutes(router, c.ProductController, env)

	return router
}

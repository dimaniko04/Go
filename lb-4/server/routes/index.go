package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.Engine, c *controllers.Controllers) *gin.RouterGroup {
	router := api.Group("/api")

	productRoutes(router, c.ProductController)

	return router
}

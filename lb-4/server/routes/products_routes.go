package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/gin-gonic/gin"
)

func productRoutes(r *gin.RouterGroup, c controllers.ProductController) {
	g := r.Group("/products")

	g.GET("/", c.GetAll)
	g.POST("/", c.Add)
	g.DELETE("/:id", c.Delete)
}

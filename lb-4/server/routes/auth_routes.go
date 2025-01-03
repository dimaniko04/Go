package routes

import (
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.RouterGroup, c controllers.AuthController) {
	r.POST("/login", c.Login)
	r.POST("/register", c.Register)
}

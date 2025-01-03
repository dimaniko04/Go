package controllers

import (
	"net/http"

	"github.com/dimaniko04/Go/lb-4/server/requests"
	"github.com/dimaniko04/Go/lb-4/server/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type authController struct {
	authService services.AuthService
}

func (c *authController) Register(ctx *gin.Context) {
	var authRequest requests.AuthRequest
	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := c.authService.Register(authRequest.Email, authRequest.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, authResponse)
}

func (c *authController) Login(ctx *gin.Context) {
	var authRequest requests.AuthRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := c.authService.Login(authRequest.Email, authRequest.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, authResponse)
}

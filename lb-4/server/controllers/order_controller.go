package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimaniko04/Go/lb-4/server/models"
	"github.com/dimaniko04/Go/lb-4/server/services"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetUserOrders(*gin.Context)
}

type orderController struct {
	orderService services.OrderService
}

func (c *orderController) GetUserOrders(ctx *gin.Context) {
	userId := ctx.Keys["user"].(models.User).Id
	orders, err := c.orderService.GetUserOrders(strconv.Itoa(userId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

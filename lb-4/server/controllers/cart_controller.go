package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dimaniko04/Go/lb-4/server/models"
	"github.com/dimaniko04/Go/lb-4/server/requests"
	"github.com/dimaniko04/Go/lb-4/server/services"
	"github.com/gin-gonic/gin"
)

type CartController interface {
	GetAll(*gin.Context)
	Add(*gin.Context)
	Delete(*gin.Context)
	Checkout(*gin.Context)
}

type cartController struct {
	cartService     services.CartService
	paymentService  services.PaymentService
	liqpayPublicKey string
}

func (c *cartController) GetAll(ctx *gin.Context) {
	userId := ctx.Keys["user"].(models.User).Id
	cartItems, err := c.cartService.GetAll(strconv.Itoa(userId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cartItems)
}

func (c *cartController) Add(ctx *gin.Context) {
	var addItemRequest requests.AddCartItemRequest
	userId := ctx.Keys["user"].(models.User).Id
	if err := ctx.ShouldBindJSON(&addItemRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.CartItem{
		UserId:    userId,
		Quantity:  addItemRequest.Quantity,
		ProductId: addItemRequest.ProductId,
	}

	err := c.cartService.Add(item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

func (c *cartController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.cartService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product removed from cart!"})
}

func (c *cartController) Checkout(ctx *gin.Context) {
	userId := ctx.Keys["user"].(models.User).Id
	var checkoutRequest requests.CheckoutRequest

	if err := ctx.ShouldBindJSON(&checkoutRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItems, err := c.cartService.GetAll(strconv.Itoa(userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parts := strings.Split(checkoutRequest.CardExpiryDate, "/")
	if len(parts) != 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid card expiry date format"})
		return
	}

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	request := requests.PaymentRequest{
		PublicKey:    c.liqpayPublicKey,
		Version:      "3",
		Action:       "pay",
		Amount:       fmt.Sprintf("%.2f", totalPrice),
		Currency:     "UAH",
		Description:  checkoutRequest.Description,
		Card:         checkoutRequest.Card,
		CardExpYear:  parts[1],
		CardExpMonth: parts[0],
		CardCvv:      checkoutRequest.CardCvv,
	}

	err = c.paymentService.Payment(request)
	if err != nil {
		ctx.JSON(http.StatusPaymentRequired, gin.H{"error": err.Error()})
		return
	}

	err = c.cartService.Checkout(strconv.Itoa(userId))
	if err != nil {
		ctx.JSON(http.StatusPaymentRequired, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Checkout successful"})
}

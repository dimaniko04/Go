package controllers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/dimaniko04/Go/lb-4/server/requests"
	"github.com/dimaniko04/Go/lb-4/server/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAll(*gin.Context)
	Add(*gin.Context)
	Delete(*gin.Context)
}

type productController struct {
	productService services.ProductService
}

func (c *productController) GetAll(ctx *gin.Context) {
	products, err := c.productService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *productController) Add(ctx *gin.Context) {
	var product requests.AddProduct

	if err := ctx.Bind(&product); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	img, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	imagePath := filepath.Join("static", img.Filename)
	if err := ctx.SaveUploadedFile(img, imagePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.productService.Add(product, imagePath)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Product added successfully!"})
}

func (c *productController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.productService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully!"})
}

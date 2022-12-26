package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

type productController struct {
	productService product.Service
}

func NewProductController(router *gin.Engine, service product.Service) {
	ctrl := &productController{
		productService: service,
	}
	router.GET("/products", ctrl.FindAll)
	router.POST("/products", ctrl.Create)
}

func (ctrl *productController) FindAll(c *gin.Context) {

	var products []schema.Product
	products, err := ctrl.productService.FindAll()

	if err != nil {
		panic("Failed to retrieve all products: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (ctrl *productController) Create(c *gin.Context) {
	// Validate input
	var input schema.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := ctrl.productService.Create(input)

	if err != nil {
		panic("Failed to create Product: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

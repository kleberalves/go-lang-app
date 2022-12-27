package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
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

	var items []schema.Product
	items, err := ctrl.productService.FindAll()

	if err != nil {
		panic("Failed to retrieve all products: " + err.Error())
	}

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *productController) Create(c *gin.Context) {
	// Validate input
	var input schema.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := ctrl.productService.Create(input)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service product.Service
}

func NewProductController(router *gin.Engine, service product.Service) {
	ctrl := &controller{
		service: service,
	}
	router.GET("/products", ctrl.FindAll)
	router.POST("/products", ctrl.Create)
	router.PUT("/products", ctrl.Update)
	router.DELETE("/products", ctrl.Delete)
}

func (ctrl *controller) FindAll(c *gin.Context) {

	var items []schema.Product
	items, err := ctrl.service.FindAll()

	if err != nil {
		panic("Failed to retrieve all products: " + err.Error())
	}

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) Create(c *gin.Context) {
	// Validate input
	var input schema.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := ctrl.service.Create(input)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) Update(c *gin.Context) {
	// Validate input
	var input schema.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.Update(input)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err})
}

func (ctrl *controller) Delete(c *gin.Context) {
	// Validate input
	var itemIds []int
	if err := c.ShouldBindJSON(&itemIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.Delete(itemIds)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err})
}

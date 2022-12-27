package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type purchaseController struct {
	purchaseService purchase.Service
}

func NewPurchaseController(router *gin.Engine, service purchase.Service) {
	ctrl := &purchaseController{
		purchaseService: service,
	}
	router.GET("/purchases", ctrl.FindAll)
	router.POST("/purchases", ctrl.Create)

}

func (ctrl *purchaseController) FindAll(c *gin.Context) {

	var items []schema.Purchase
	items, err := ctrl.purchaseService.FindAll()

	if err != nil {
		panic("Failed to retrieve all purchases: " + err.Error())
	}

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *purchaseController) Create(c *gin.Context) {
	// Validate input
	var input schema.Purchase
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := ctrl.purchaseService.Create(input)

	if err != nil {
		panic("Failed to create User: " + err.Error())
	}

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

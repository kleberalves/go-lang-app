package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service purchase.Service
}

func NewPurchaseController(router *gin.Engine, service purchase.Service) {
	ctrl := &controller{
		service: service,
	}
	router.GET("/purchases", ctrl.FindAll)
	router.GET("/purchases/:userid", ctrl.GetByUser)
	router.POST("/purchases", ctrl.Create)
	router.DELETE("/purchases", ctrl.Delete)

}

func (ctrl *controller) FindAll(c *gin.Context) {

	var items []schema.Purchase
	items, err := ctrl.service.FindAll()

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) GetByUser(c *gin.Context) {

	paramUserId := c.Param("userid")
	userId, err := strconv.Atoi(paramUserId)

	if err != nil {
		panic("Failed to convert USERID parameter: " + err.Error())
	}

	items, err := ctrl.service.GetByUser(userId)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) Create(c *gin.Context) {
	// Validate input
	var input schema.Purchase
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

func (ctrl *controller) Delete(c *gin.Context) {

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

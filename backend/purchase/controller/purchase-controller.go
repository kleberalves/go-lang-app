package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/purchase"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service    purchase.Service
	credential credential.Service
}

func NewPurchaseController(router *gin.Engine, service purchase.Service, credential credential.Service) {
	ctrl := &controller{
		service:    service,
		credential: credential,
	}

	onlySalesman := router.Group("/purchases")
	onlySalesman.Use(services.JwtAuthMiddlewareRoles(credential,
		[]enums.TypeUser{enums.Salesman}))
	onlySalesman.GET("/", ctrl.FindAll)
	onlySalesman.POST("/salesman", ctrl.CreateByAuthenticatedSalesman)
	onlySalesman.DELETE("/", ctrl.Delete)

	onlyCustomer := router.Group("/purchases")
	onlyCustomer.Use(services.JwtAuthMiddlewareRoles(credential,
		[]enums.TypeUser{enums.Customer}))
	onlyCustomer.POST("/customer", ctrl.CreateByAuthenticatedCustomer)
	onlyCustomer.GET("/my", ctrl.GetByMyId)

}

func (ctrl *controller) FindAll(c *gin.Context) {

	var items []schema.Purchase
	items, err := ctrl.service.FindAll()

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) GetByMyId(c *gin.Context) {

	items, err := ctrl.service.GetByMyId(c)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) CreateByAuthenticatedSalesman(c *gin.Context) {
	// Validate input
	var input schema.Purchase
	if !httphandler.GetJson(&input, c) {
		return
	}

	item, err := ctrl.service.CreateByAuthenticatedSalesman(input, c)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) CreateByAuthenticatedCustomer(c *gin.Context) {
	// Validate input
	var input schema.Purchase
	if !httphandler.GetJson(&input, c) {
		return
	}

	item, err := ctrl.service.CreateMyPurchase(input, c)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) Delete(c *gin.Context) {

	var input []int
	if !httphandler.GetJson(&input, c) {
		return
	}

	err := ctrl.service.Delete(input)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err})
}

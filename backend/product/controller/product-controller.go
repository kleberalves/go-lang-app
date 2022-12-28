package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/product"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service product.Service
}

func NewProductController(router *gin.Engine, service product.Service, credential credential.Service) {
	ctrl := &controller{
		service: service,
	}

	onlySalesman := router.Group("/products")
	onlySalesman.Use(services.JwtAuthMiddlewareRoles(credential,
		[]enums.TypeUser{enums.Salesman}))
	onlySalesman.POST("/", ctrl.Create)
	onlySalesman.PUT("/", ctrl.Update)
	onlySalesman.DELETE("/", ctrl.Delete)

	justValidToken := router.Group("/products")
	justValidToken.Use(services.JwtAuthMiddleware())
	justValidToken.GET("/", ctrl.FindAll)
}

func (ctrl *controller) FindAll(c *gin.Context) {

	var items []schema.Product
	items, err := ctrl.service.FindAll()

	if err != nil {
		panic("Failed to retrieve all products: " + err.Error())
	}

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) Create(c *gin.Context) {

	var input schema.Product
	if !httphandler.GetJson(&input, c) {
		return
	}

	item, err := ctrl.service.Create(input)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) Update(c *gin.Context) {
	var input schema.Product
	if !httphandler.GetJson(&input, c) {
		return
	}

	err := ctrl.service.Update(input)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err})
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

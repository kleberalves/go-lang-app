package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
	"github.com/kleberalves/problemCompanyApp/backend/user"
	"github.com/kleberalves/problemCompanyApp/backend/user/filter"
)

type controller struct {
	service user.Service
}

func NewUserController(router *gin.Engine, service user.Service, credential credential.Service) {
	ctrl := &controller{
		service: service,
	}

	onlySalesman := router.Group("/users")
	onlySalesman.Use(services.JwtAuthMiddlewareRoles(credential,
		[]enums.TypeUser{enums.Salesman}))
	onlySalesman.GET("/", ctrl.FindAll)
	onlySalesman.POST("/", ctrl.Create)
	onlySalesman.GET("/:id", ctrl.Get)
	onlySalesman.PUT("/", ctrl.Update)
	onlySalesman.DELETE("/", ctrl.Delete)

	justValidToken := router.Group("/users")
	justValidToken.Use(services.JwtAuthMiddleware())
	justValidToken.GET("/salesman/:name", ctrl.FindSalesmanByName)

}

func (ctrl *controller) FindAll(c *gin.Context) {

	var input filter.UserFilter
	if !httphandler.GetJson(&input, c) {
		return
	}

	var items []schema.UserRead
	items, err := ctrl.service.FindByFilter(input)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) FindSalesmanByName(c *gin.Context) {

	name := c.Param("name")

	filter := filter.UserFilter{
		FirstName:   name,
		ProfileType: 1,
	}

	var items []schema.UserRead
	items, err := ctrl.service.FindByFilter(filter)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) Create(c *gin.Context) {

	var input schema.User
	if !httphandler.GetJson(&input, c) {
		return
	}

	item, err := ctrl.service.Create(input)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) Get(c *gin.Context) {

	paramdId := c.Param("id")
	id, err := strconv.Atoi(paramdId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := ctrl.service.Get(id)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) Update(c *gin.Context) {

	var input schema.User
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

	if err != nil {
		panic("Failed to delete Users: " + err.Error())
	}

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err})
}

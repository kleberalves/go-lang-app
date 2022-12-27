package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
	"github.com/kleberalves/problemCompanyApp/backend/user"
)

type controller struct {
	service user.Service
}

func NewUserController(router *gin.Engine, service user.Service) {
	ctrl := &controller{
		service: service,
	}
	router.GET("/users", ctrl.FindAll)
	router.POST("/users", ctrl.Create)
	router.GET("/users/:id", ctrl.Get)
	router.PUT("/users", ctrl.Update)
	router.DELETE("/users", ctrl.Delete)

	router.POST("/users/login", ctrl.Create)

}

func (ctrl *controller) FindAll(c *gin.Context) {

	var items []schema.UserRead
	items, err := ctrl.service.FindAll()

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) Create(c *gin.Context) {
	// Validate input
	var input schema.User
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

func (ctrl *controller) Get(c *gin.Context) {

	paramdId := c.Param("id")
	id, err := strconv.Atoi(paramdId)

	if err != nil {
		panic("Failed to convert ID parameter: " + err.Error())
	}

	item, err := ctrl.service.Get(id)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) Update(c *gin.Context) {
	// Validate input
	var input schema.User
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
	var input []int
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.Delete(input)

	if err != nil {
		panic("Failed to delete Users: " + err.Error())
	}

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err})
}

package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/profile"
	"github.com/kleberalves/problemCompanyApp/backend/services"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service profile.Service
}

func NewProfileController(router *gin.Engine, service profile.Service, credential credential.Service) {
	ctrl := &controller{
		service: service,
	}

	onlySalesman := router.Group("/profiles")
	onlySalesman.Use(services.JwtAuthMiddlewareRoles(credential,
		[]enums.TypeUser{enums.Salesman}))

	onlySalesman.GET("/", ctrl.FindAll)
	onlySalesman.POST("/:typoid/user/:id", ctrl.AddProfile)
	onlySalesman.DELETE("/:typoid", ctrl.RemoveProfiles)
}

func (ctrl *controller) FindAll(c *gin.Context) {

	items, err := ctrl.service.FindAll()

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     items})
}

func (ctrl *controller) AddProfile(c *gin.Context) {

	paramdId := c.Param("id")
	id, err := strconv.Atoi(paramdId)

	if err != nil {
		panic("Failed to convert ID parameter: " + err.Error())
	}

	paramTypo := c.Param("typoid")
	typo, err2 := strconv.Atoi(paramTypo)

	if err2 != nil {
		panic("Failed to convert TYPO parameter: " + err2.Error())
	}

	item, err := ctrl.service.AddProfile(id, typo)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

func (ctrl *controller) RemoveProfiles(c *gin.Context) {

	paramTypo := c.Param("typoid")
	typo, err := strconv.Atoi(paramTypo)

	if err != nil {
		panic("Failed to convert TYPO parameter: " + err.Error())
	}

	var input []int
	if !httphandler.GetJson(&input, c) {
		return
	}

	err = ctrl.service.RemoveProfiles(input, typo)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err})
}

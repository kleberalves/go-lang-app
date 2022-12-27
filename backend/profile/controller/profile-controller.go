package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/profile"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service profile.Service
}

func NewProfileController(router *gin.Engine, service profile.Service) {
	ctrl := &controller{
		service: service,
	}
	router.GET("/profiles", ctrl.FindAll)
	router.POST("/profiles/:typoid/user/:id", ctrl.AddProfile)
	router.DELETE("/profiles/:typoid", ctrl.RemoveProfiles)
}

func (ctrl *controller) FindAll(c *gin.Context) {

	items, err := ctrl.service.FindAll()

	httphandler.Response(httphandler.RParams{
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

	httphandler.Response(httphandler.RParams{
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

	var userIds []int
	if err := c.ShouldBindJSON(&userIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.RemoveProfiles(userIds, typo)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err})
}

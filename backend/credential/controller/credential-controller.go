package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	httphandler "github.com/kleberalves/problemCompanyApp/backend/services/http-handler"
)

type controller struct {
	service credential.Service
}

func NewCredentialController(router *gin.Engine, service credential.Service) {
	ctrl := &controller{
		service: service,
	}
	router.POST("/credential/login", ctrl.Login)
	/*
		FUTURE

		Increase security with recaptcha on anonymous calls
		router.POST("/credential/login/:recaptcha", ctrl.Login)
		router.POST("/credential/login/activate/:recaptcha", ctrl.XXX)
		router.POST("/credential/reset-password-send/:email/:recaptcha", ctrl.XXX)
		router.POST("/credential/reset-password/:recaptcha", ctrl.XXX)

		Enable - On Time Password
		router.GET("/credential/generate-secret-otp", ctrl.XXX)
		router.POST("/credential/validate-otp", ctrl.XXX)
		router.POST("/credential/remove-otp", ctrl.XXX)

	*/
}

func (ctrl *controller) Login(c *gin.Context) {
	var input schema.UserCredential
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := ctrl.service.Login(input)

	httphandler.Response(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

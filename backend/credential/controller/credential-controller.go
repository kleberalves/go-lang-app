package controller

import (
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

}

func (ctrl *controller) Login(c *gin.Context) {
	var input schema.UserCredential
	if !httphandler.GetJson(&input, c) {
		return
	}

	item, err := ctrl.service.Login(input)

	httphandler.ResponseCheck(httphandler.RParams{
		Context: c,
		Err:     err,
		Obj:     item})
}

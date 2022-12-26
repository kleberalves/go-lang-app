package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/user"
)

type userController struct {
	userService user.Service
}

func NewUserController(router *gin.Engine, service user.Service) {
	ctrl := &userController{
		userService: service,
	}
	router.GET("/users", ctrl.FindAll)
	router.POST("/users", ctrl.Create)
	router.POST("/users/:id/profiles/:typoid", ctrl.AssociateProfile)
	router.DELETE("/users/:id/profiles/:typoid", ctrl.RemoveProfile)
	// router.GET("/articles/:id", handler.GetByID)
	// router.DELETE("/articles/:id", handler.Delete)
}

func (ctrl *userController) AssociateProfile(c *gin.Context) {
	// Validate input

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

	profile := ctrl.userService.AssociateProfile(id, typo)

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

func (ctrl *userController) RemoveProfile(c *gin.Context) {
	// Validate input

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

	ctrl.userService.RemoveProfile(id, typo)

	c.Writer.WriteHeader(http.StatusNoContent)
}

func (ctrl *userController) FindAll(c *gin.Context) {

	var users []schema.UserRead
	users, err := ctrl.userService.FindAll()

	if err != nil {
		panic("Failed to retrieve all Users: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (ctrl *userController) Create(c *gin.Context) {
	// Validate input
	var input schema.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.Create(input)

	if err != nil {
		panic("Failed to create User: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

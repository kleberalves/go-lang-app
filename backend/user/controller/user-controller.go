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
	router.GET("/users/:id", ctrl.Get)
	router.PUT("/users", ctrl.Update)
	router.DELETE("/users", ctrl.Delete)
	router.POST("/users/:id/profiles/:typoid", ctrl.AssociateProfile)
	router.DELETE("/users/:id/profiles/:typoid", ctrl.RemoveProfile)
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

	profile := ctrl.userService.AddProfile(id, typo)

	c.JSON(http.StatusOK, profile)
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

	ctrl.userService.DeleteProfile(id, typo)

	c.Writer.WriteHeader(http.StatusNoContent)
}

func (ctrl *userController) FindAll(c *gin.Context) {

	var users []schema.UserRead
	users, err := ctrl.userService.FindAll()

	if err != nil {
		panic("Failed to retrieve all Users: " + err.Error())
	}

	c.JSON(http.StatusOK, users)
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

	c.JSON(http.StatusOK, user)
}

func (ctrl *userController) Get(c *gin.Context) {

	paramdId := c.Param("id")
	id, err := strconv.Atoi(paramdId)

	if err != nil {
		panic("Failed to convert ID parameter: " + err.Error())
	}

	user, err := ctrl.userService.Get(id)

	if err != nil {
		panic("Failed to get User: " + err.Error())
	}

	c.JSON(http.StatusOK, user)

}

func (ctrl *userController) Update(c *gin.Context) {
	// Validate input
	var input schema.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.userService.Update(input)

	if err != nil {
		panic("Failed to update User: " + err.Error())
	}

	c.Writer.WriteHeader(http.StatusNoContent)

}

func (ctrl *userController) Delete(c *gin.Context) {
	// Validate input
	var input []int
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.userService.Delete(input)

	if err != nil {
		panic("Failed to delete Users: " + err.Error())
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

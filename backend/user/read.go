package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
)

// GET /users
// Get all users
func Read(c *gin.Context) {

	db := schema.Connect()

	var users []schema.UserRead
	err := db.Model(&schema.User{}).Preload("Profiles").Find(&users).Error

	if err != nil {
		panic("Failed to retrieve all Users: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

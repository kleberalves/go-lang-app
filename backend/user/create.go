package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/schema"
	"github.com/kleberalves/problemCompanyApp/backend/services"
)

func Create(c *gin.Context) {
	// Validate input
	var input schema.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := services.HashPassword(input.Password)

	user := schema.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hash,
		Profiles:  input.Profiles}

	db := schema.Connect()
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

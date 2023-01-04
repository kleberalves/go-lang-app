package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/services/security"
)

type ResponseError struct {
	Message string `json:"message"`
}

func JwtAuthMiddlewareRoles(credential credential.Service, r []enums.TypeUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := security.TokenValid(c)
		if err != nil {
			finishUnauthorizedToken(c)
			return
		}
		if len(r) > 0 {
			id, err := security.ExtractTokenID(c)
			if err != nil {
				finishUnauthorized(c)
				return
			}

			rolesOk, err := credential.CheckRolesByUserID(id, r)

			if !rolesOk || err != nil {
				finishUnauthorized(c)
				return
			}
		}

		c.Next()
	}
}
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := security.TokenValid(c)
		if err != nil {
			finishUnauthorizedToken(c)
			return
		}
		c.Next()
	}
}

func finishUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ResponseError{Message: "Unauthorized"})
	c.Abort()
}

func finishUnauthorizedToken(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, ResponseError{Message: "Unauthorized token"})
	c.Abort()
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//TODO: Update env variable to set production origin URL
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("ACCESS_ORIGIN_URL"))
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

		if c.Request.Method == http.MethodOptions {

			fmt.Println(c.GetHeader("Access-Control-Request-Method"))
			c.Writer.WriteHeader(http.StatusOK)
			return
		}
	}
}

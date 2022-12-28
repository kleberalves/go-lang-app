package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kleberalves/problemCompanyApp/backend/credential"
	"github.com/kleberalves/problemCompanyApp/backend/enums"
	"github.com/kleberalves/problemCompanyApp/backend/services/security"
)

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
	c.Writer.WriteHeader(http.StatusUnauthorized)
	c.Abort()
}

func finishUnauthorizedToken(c *gin.Context) {
	c.String(http.StatusUnauthorized, "Unauthorized token")
	c.Abort()
}

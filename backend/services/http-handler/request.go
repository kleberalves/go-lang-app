package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_cors "github.com/rs/cors"
)

func GetJson(obj any, c *gin.Context) bool {

	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	return true
}

func CorsMiddleware(optionPassthrough bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//TODO: create env variable

		cors := _cors.New(_cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000", "/api"},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production

			Debug: true,
		})

		cors.HandlerFunc(ctx.Writer, ctx.Request)
		if !optionPassthrough &&
			ctx.Request.Method == http.MethodOptions &&
			ctx.GetHeader("Access-Control-Request-Method") != "" {
			// Abort processing next Gin middlewares.
			ctx.AbortWithStatus(http.StatusOK)
		}
	}
}

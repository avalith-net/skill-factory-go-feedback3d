package middleware

import (
	"net/http"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/controller"
	"github.com/gin-gonic/gin"
)

//ValidateJWT is used to check the jwt passed as parameter.
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, _, err := controller.TokenProcess(c.GetHeader("Authorization"))
		if err != nil {
			c.String(http.StatusUnauthorized, "Must log in.")
			c.Abort()
		}
		c.Next()
	}
}

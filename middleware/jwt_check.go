package middleware

import (
	"net/http"

	"github.com/blotin1993/feedback-api/controller"
	"github.com/gin-gonic/gin"
)

//ValidateJWT is used to check the jwt passed as parameter.
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, _, err := controller.TokenProcess(c.GetHeader("Authorization"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Token error.")
			return
		}
		c.Next()
	}
}

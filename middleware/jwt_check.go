package middleware

import (
	"github.com/blotin1993/feedback-api/routers"
	"github.com/gin-gonic/gin"
)

//ValidateJWT is used to check the jwt passed as parameter.
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, _, err := routers.TokenProcess(c.GetHeader("Authorization"))
		if err != nil {
			// http.Error(w, "Token error."+err.Error(), http.StatusBadRequest)
			c.AbortWithStatusJSON(500, gin.H{"message": "Token error."})
			return
		}
		c.Next()
	}
}

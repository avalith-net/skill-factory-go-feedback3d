package middleware

import (
	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

//CheckDb middleware.
func CheckDb() gin.HandlerFunc {
	return func(c *gin.Context) {
		if db.CheckConnection() == 0 {
			c.AbortWithStatusJSON(500, gin.H{"message": "Connection lost."})
			// http.Error(w, "Connection lost.", 500)
			return
		}
		c.Next()
	}
}

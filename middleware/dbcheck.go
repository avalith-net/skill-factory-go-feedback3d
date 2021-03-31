package middleware

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

//CheckDb middleware.
func CheckDb() gin.HandlerFunc {
	return func(c *gin.Context) {
		if db.CheckConnection() == 0 {
			c.String(http.StatusInternalServerError, "Connection lost.")
			c.Abort()
		}
		c.Next()
	}
}

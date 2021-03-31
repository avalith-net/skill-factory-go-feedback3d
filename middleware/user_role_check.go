package middleware

import (
	"net/http"

	"github.com/blotin1993/feedback-api/controller"
	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

//userRoleCheck middleware.
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := db.GetUser(controller.IDUser)
		if err != nil {
			c.String(http.StatusInternalServerError, "User error.")
			c.Abort()
		}
		if user.Role != "admin" {
			c.String(http.StatusUnauthorized, "Must be admin.")
			c.Abort()
		}
		c.Next()
	}
}

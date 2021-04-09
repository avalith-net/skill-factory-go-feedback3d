package middleware

import (
	"net/http"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/controller"
	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
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

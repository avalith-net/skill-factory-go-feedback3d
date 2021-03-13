package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

//ChangePassEmail .
func ChangePassEmail(c *gin.Context) {
	//parámetros: id, token, newpass
	id := c.Query("id")
	token := c.Query("token")
	newPass := c.Query("newpass")
	if len(newPass) < 6 {
		c.String(http.StatusBadRequest, "The new password must be at least 6 characters long.")
		return
	}
	user := models.User{
		Password: newPass,
	}
	_, isOk, _, _ := TokenProcess(token)
	if !isOk {
		c.String(http.StatusBadRequest, "Authentication error.")
		return
	}
	//modificar usuario
	hasEffect, err := db.ModifyUser(user, id)
	if !hasEffect {
		c.String(http.StatusInternalServerError, "An error has ocurred trying to set a new password."+err.Error())
		return
	}
	c.String(http.StatusCreated, "Success")
}

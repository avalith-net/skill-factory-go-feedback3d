package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

func BanUser(c *gin.Context) {
	email := c.Param("email")

	user, found, id := db.UserAlreadyExist(email)
	if user.Enabled == false {
		c.String(http.StatusBadRequest, "User already banned.")
		return
	}
	if !found {
		c.String(http.StatusBadRequest, "User does not exist.")
		return
	}
	var userEmpty models.User
	userEmpty.Enabled = false

	_, err := db.ModifyUser(userEmpty, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error trying to modify the user")
		return
	}

	c.String(http.StatusCreated, "User banned.")
}

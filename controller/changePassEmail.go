package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

// ChangePassEmail godoc
// @Description get string by id, token and newpass
// @id changepass
// @Summary is used to handle the mail you get when recovering your password.
// @Param id query string true "Account ID"
// @Param token query string true "JWT Token"
// @Param newpass query string true "New Pass"
// @Produce plain
// @Success 201 {string} string "Password has been changed."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Validation error"
// @Failure 500 {string} string "An error has ocurred trying to set a new password."
// @Failure default {string} string "An error has ocurred"
// @Router /changePassword [post]
func ChangePassEmail(c *gin.Context) {

	//init pass variable to hold json incoming data.
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Check your form.")
		return
	}
	if len(user.Password) < 6 {
		c.String(http.StatusBadRequest, "The new password must be at least 6 characters long.")
		return
	}

	//modificar usuario
	hasEffect, err := db.ModifyUser(user, IDUser)
	if !hasEffect {
		c.String(http.StatusInternalServerError, "An error has ocurred trying to set a new password."+err.Error())
		return
	}
	c.String(http.StatusCreated, "Success")
}

package controller

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/auth"
	"github.com/blotin1993/feedback-api/db"
	services "github.com/blotin1993/feedback-api/services/email"
	"github.com/gin-gonic/gin"
)

// RecoverPass godoc
// @Description receive the user data from DataBase and send an Email with his current password
// @id recpass
// @Summary is used to recover our password
// @Param email query string true "email"
// @Success 201 {string} string "Email sent."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Wrong mail."
// @Failure 500 {string} string "Internal Error."
// @Failure default {string} string "An error has ocurred"
// @Router /recoverPass [post]
func RecoverPass(c *gin.Context) {
	email := c.Query("email")
	if len(email) < 1 {
		c.String(http.StatusBadRequest, "must complete email form")
		return
	}
	user, mailExist, _ := db.UserAlreadyExist(email)
	if !mailExist {
		c.String(http.StatusBadRequest, "Wrong mail.")
		return
	}

	stringObjectID := (user.ID).Hex()
	expirationTime := time.Now().Add(1 * time.Hour)
	jwtKey, err := auth.GenerateJWT(user, expirationTime)

	bodyString := "Hey <b><i>" + user.Name + "</i></b>!\nFollow this link to recover your password.\n <b><i>http:localhost:8080/changePassword?id=" + stringObjectID + "&token=Bearer " + jwtKey + "</i></b>"

	//Email send function
	if !services.SendEmail(email, "Get your password.", bodyString) {
		c.String(http.StatusBadRequest, "An error has ocurred sending the email"+err.Error())
		return
	}
	c.String(http.StatusCreated, "Success")
}

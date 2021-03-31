package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

// getGeneralProf godoc
// @Description get string by id
// @id getGeneralProf
// @Summary is used to request a feedback to other user.
// @Param id query string true "Account ID"
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /fbRequest [post]
func GetGeneralProfile(c *gin.Context) {

	var userGeneral models.GeneralProfile
	var fullName string

	id := c.Query("id")
	fullName = c.Query("fullname")
	if len(fullName) < 1 {
		c.String(http.StatusBadRequest, "You must enter a name.")
		return
	}

	issuer, receiver, err := db.GetFeedbacks(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Internal error.")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exist.")
		return
	}
	profilePic, err := db.GetProfilePic(user.Email)
	if err != nil {
		c.String(http.StatusBadRequest, "Picture error.")
		return
	}

	userGeneral.CompleteName = user.Name + user.LastName
	userGeneral.ProfilePicture = profilePic
	userGeneral.FbIssuer = len(issuer)
	userGeneral.FbReceiver = len(receiver)

	c.JSON(http.StatusCreated, userGeneral)
}

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
// @Summary is used to get the profile of some user.
// @Param id path string true "Account ID"
// @Param Authorization header string true "jwt token"
// @Produce json
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /users/get/{id} [get]
func GetGeneralProfile(c *gin.Context) {

	id := c.Param("id")

	if len(id) < 1 {
		c.String(http.StatusBadRequest, "User ID is needed.")
		return
	}

	receiver, err := db.GetFeedFromDb(id, true)
	issuer, err2 := db.GetFeedFromDb(id, false)

	if err != nil || err2 != nil {
		c.String(http.StatusInternalServerError, "Internal error.")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exist.")
		return
	}

	var userGeneral models.GeneralProfile
	userGeneral.CompleteName = user.Name + " " + user.LastName
	userGeneral.ProfilePicture = user.ProfilePicture
	userGeneral.FbIssuer = len(issuer)
	userGeneral.FbReceiver = len(receiver)

	user, _ = db.GetUser(IDUser)

	if user.Role == "admin" {
		results, err := db.GetFeedFromDb(id, true)
		if err != nil {
			c.String(http.StatusInternalServerError, "Internal error.")
			return
		}
		var adminProfile models.AdminProfile
		adminProfile.Profile = userGeneral
		adminProfile.Metrics = results
		c.JSON(http.StatusCreated, adminProfile)
		return
	}

	c.JSON(http.StatusCreated, userGeneral)
}

package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

// dashboard godoc
// @Description get user feedbacks from db.
// @id dash
// @Summary is used to get feedbacks of the user.
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Success"
// @Failure 400 {string} string "Error ID."
// @Failure default {string} string "An error has ocurred"
// @Router /dashboard [get]
func GetDashboard(c *gin.Context) {
	user, err := db.GetUser(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "Error"+err.Error())
		return
	}
	feedSlice, err := db.GetFeedFromDb(IDUser, true)
	if err != nil {
		c.String(http.StatusBadRequest, "Error"+err.Error())
		return
	}
	var userGeneral models.AdminProfile
	userGeneral.Profile.CompleteName = user.Name + " " + user.LastName
	userGeneral.Profile.ProfilePicture = user.ProfilePicture
	userGeneral.Profile.Graphic = user.Graphic
	userGeneral.Metrics = feedSlice

	c.JSON(http.StatusCreated, userGeneral)
}

package controller

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
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

	//Getting all the feedRequestedObjs of the logged user.
	allFeedRequested, err := db.GetAllFeedRequested(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "There's no feedRequested objects.")
		return
	}

	allUsersWhoAskedForFeed, err := db.GetAllUsersAskingForFeed(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "There's no UsersAsksFeed objects.")
		return
	}

	var userGeneral models.AdminProfile
	userGeneral.Profile.CompleteName = user.Name + " " + user.LastName
	userGeneral.Profile.ProfilePicture = user.ProfilePicture
	userGeneral.Profile.Graphic = user.Graphic
	userGeneral.Metrics = feedSlice
	userGeneral.Profile.FeedbackSent = len(user.FeedbackStatus.FeedbacksSended)
	userGeneral.Profile.FeedbacksRequested = len(allFeedRequested)
	userGeneral.Profile.FeedbackAskedForUsers = len(allUsersWhoAskedForFeed)

	c.JSON(http.StatusCreated, userGeneral)
}

package controller

import (
	"fmt"
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
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

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exist.")
		return
	}

	//Getting all the feedRequestedObjs of the logged user.
	allFeedRequested, err := db.GetAllFeedRequested(id)
	if err != nil {
		c.String(http.StatusBadRequest, "There's no feedRequested objects.")
		return
	}

	allUsersWhoAskedForFeed, err := db.GetAllUsersAskingForFeed(id)
	if err != nil {
		c.String(http.StatusBadRequest, "There's no UsersAsksFeed objects.")
		return
	}

	var userGeneral models.GeneralProfile
	userGeneral.CompleteName = user.Name + " " + user.LastName
	userGeneral.FeedbacksRequested = len(allFeedRequested)
	userGeneral.FeedbackAskedForUsers = len(allUsersWhoAskedForFeed)
	userGeneral.FeedbackSent = len(user.FeedbackStatus.FeedbacksSended)
	userGeneral.ProfilePicture = user.ProfilePicture

	fmt.Println("Imprimiendo allFeedRequested, allUsersWhoAskedForFeed y los feedback enviados")
	fmt.Println(len(allFeedRequested))
	fmt.Println(len(allUsersWhoAskedForFeed))
	fmt.Println(len(user.FeedbackStatus.FeedbacksSended))

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

package controller

import (
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	services "github.com/avalith-net/skill-factory-go-feedback3d/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/fatih/structs"
)

// Feedback  godoc
// @Description gives feedback to the user
// @id fb
// @Summary is used to give feedback to users.
// @Param target_id query string true "Target ID"
// @Param feedback body string true "Json body with email and password"
// @Param Authorization header string true "JWT Token"
// @Accept  json
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Wrong mail or password."
// @Failure 500 {string} string "Error generating the token."
// @Failure default {string} string "An error has ocurred"
// @Router /feedback [post]
func FeedbackTry(c *gin.Context) {
	rID := c.Query("target_id")
	if len(rID) < 1 {
		c.String(http.StatusBadRequest, "ID Error")
		return
	}
	loggedUser, _ := db.GetUser(IDUser)
	if !loggedUser.Enabled {
		c.String(http.StatusUnauthorized, "User not authorized.")
		return
	}
	user, err := db.GetUser(rID)
	if err != nil {
		c.String(http.StatusBadRequest, "Error returning user")
	}
	_, isFound, _ := db.UserAlreadyExist(user.Email)
	if !isFound {
		c.String(http.StatusBadRequest, "User was not found.")
		return
	}
	validUser, _ := db.GetUser(rID)
	if !validUser.Enabled {
		c.String(http.StatusUnauthorized, "User not authorized to receive feedbacks.")
		return
	}

	var fb models.Feedback

	if err := c.ShouldBindJSON(&fb); err != nil {
		c.String(http.StatusBadRequest, "Check your form.")
		return
	}

	//-------Feedback validation------
	if structs.IsZero(fb) || !hasZeroGroup(fb.PerformanceArea, fb.TeamArea, fb.TechArea) {
		c.String(http.StatusBadRequest, "You must enter at least one complete area")
		return
	}
	//-----------------------------------Format validation

	isValid, err := services.ValidateForm(fb)
	if !isValid {
		c.String(http.StatusBadRequest, "Invalid format: "+err.Error())
		return
	}

	// graphic stats
	user.Graphic, err = services.InitGraphic(fb, user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// persist graphic.
	_, err = db.UpdateGraphic(user, rID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------

	fb.IssuerID = IDUser
	fb.IssuerName = loggedUser.Name + " " + loggedUser.LastName
	fb.ReceiverID = rID
	fb.Date = time.Now()
	fb.IsApprobed = false
	fb.IsReported = false
	fb.IsDisplayable = true

	// Send email notification

	msg := "Hi " + user.Name + " " + user.LastName + "! \n You have received a new feedback, check it in your dashboard! <a>http:localhost:8080/dashboard</a> \n\n"
	go services.SendEmail(user.Email, "New Feedback Received!", msg)

	//------------------------------

	feedID, status, err := db.AddFeedback(fb)

	if err != nil {
		c.String(http.StatusInternalServerError, "An error has ocurred. Try again later "+err.Error())
		return
	}

	if !status {
		c.String(http.StatusInternalServerError, "Database error.")
		return
	}

	//If the feedback was requested, then inside feedback_request function the feedRequested and userAskingFeed objects should have been created.
	//if they have not been created, then the feedback was not requested and it's not necessary delete them.
	feedRequestedID, isFound := db.GetFeedBackRequestedID(IDUser, rID)
	if isFound {
		userAskingFeedID, err := db.GetUsersAskedFeedbackID(IDUser, rID)
		if err != nil {
			c.String(http.StatusBadRequest, "Error trying to get userAsksFeed ID with givens users IDs. ")
			return
		}

		_, isDeleted := db.DeleteFeedbackRequested(feedRequestedID)
		if !isDeleted {
			c.String(http.StatusBadRequest, "Error trying to delete requested feed with given ID. ")
			return
		}

		_, isDeleted = db.DeleteUserAskFeedback(userAskingFeedID)
		if !isDeleted {
			c.String(http.StatusBadRequest, "Error trying to delete asked request with given ID. ")
			return
		}
	}

	var modifiedLoggedUser models.User

	copier.Copy(&modifiedLoggedUser, &loggedUser)

	modifiedLoggedUser.FeedbackStatus.FeedbacksSended = append(modifiedLoggedUser.FeedbackStatus.FeedbacksSended, feedID)

	isLoggedUserModified, err := db.ModifyUser(modifiedLoggedUser, IDUser)
	if !isLoggedUserModified {
		c.String(http.StatusBadRequest, "Fail on changing the logged user. ")
		return
	}
	if err != nil {
		c.String(http.StatusBadRequest, "Error trying to modified logged user. ")
		return
	}

	c.String(http.StatusCreated, "Success")
}

func hasZeroGroup(group ...interface{}) bool {
	count := 0
	for _, field := range group {
		if !structs.HasZero(field) {
			return true
		}
		count++
	}
	return count != len(group)
}

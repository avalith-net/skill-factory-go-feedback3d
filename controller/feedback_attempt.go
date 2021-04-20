package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	services "github.com/avalith-net/skill-factory-go-feedback3d/services/email"
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

	/* Métricas de Feedback:
	Let´s work on this.
	Reach the Goal.
	Relevant Performance.
	Master. */
	loggedUser, _ := db.GetUser(IDUser)
	if !loggedUser.Enabled {
		c.String(http.StatusUnauthorized, "User not authorized.")
		return
	}

	rID := c.Query("target_id")
	if len(rID) < 1 {
		c.String(http.StatusBadRequest, "ID Error")
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
	err = json.NewDecoder(c.Request.Body).Decode(&fb)
	if err != nil {
		c.String(http.StatusBadRequest, "Check your form.")
		return
	}

	//-------Feedback validation------
	if structs.IsZero(fb) || !hasZeroGroup(fb.PerformanceArea, fb.TeamArea, fb.TechArea) {
		c.String(http.StatusBadRequest, "You must enter at least one complete area")
		return
	}

	if !validateMsgLength(1615, fb.Message) { //1615 xq toma saltos de pagina como caracteres.
		c.String(http.StatusBadRequest, "Message cannot be longer than 1500 characters")
		return
	}
	if !validateMsgLength(541, fb.TechArea.Message, fb.TeamArea.Message, fb.PerformanceArea.Message) { //40 xq toma saltos de página
		c.String(http.StatusBadRequest, "Area Messages cannot be longer than 500 characters.")
		return
	}
	//-----------------------------------

	fb.IssuerID = IDUser
	fb.ReceiverID = rID
	fb.Date = time.Now()
	fb.IsApprobed = false
	fb.IsReported = false

	// Send email notification

	msg := "Hi " + user.Name + " " + user.LastName + "! \n You have received a new feedback, check it in your dashboard! <a>http:localhost:8080/dashboard</a> \n\n"
	services.SendEmail(user.Email, "New Feedback Received!", msg)

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

	feedRequestedID, err := db.GetFeedBackRequestedID(IDUser, rID)
	if err != nil {
		c.String(http.StatusBadRequest, "Error trying to get feedbackRequest ID with givens users IDs. ")
		return
	}

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

	//----------------------------------------------------------------------------------------------------

	c.String(http.StatusCreated, "Success")
}

func validateMsgLength(maxLen int, Amsg ...string) bool {
	for _, msg := range Amsg {
		if len(msg) > maxLen {
			return false
		}
	}
	return true
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

package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	services "github.com/JoaoPaulo87/skill-factory-go-feedback3d/services/email"
	"github.com/gin-gonic/gin"

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
	validUser, _ := db.GetUser(IDUser)
	if !validUser.Enabled {
		c.String(http.StatusUnauthorized, "User not authorized.")
		return
	}

	rID := c.Query("target_id")
	if len(rID) < 1 {
		c.String(http.StatusBadRequest, "ID Error")
		return
	}
	user, err := db.GetUser(rID)
	_, isFound, _ := db.UserAlreadyExist(user.Email)
	if !isFound {
		c.String(http.StatusBadRequest, "User was not found.")
		return
	}
	validUser, _ = db.GetUser(rID)
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

	_, status, err := db.AddFeedback(fb)

	if err != nil {
		c.String(http.StatusInternalServerError, "An error has ocurred. Try again later "+err.Error())
		return
	}

	if !status {
		c.String(http.StatusInternalServerError, "Database error.")
		return
	}

	// // We search for the feedback just created inside the db so we can get it back with the given objID.
	// feed, err := db.GetSelectedFeedBack(feedbackID)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "The feedback with given ID does not exists.")
	// 	return
	// }

	//var modifiedUser models.User

	// Feedback is added to the FeedbackStatus of the target user, so when you are looking at their profile,
	// all pending feedbacks are displayed. This is connected in get_generalProfile
	//modifiedUser.UsersAskedFeed = append(modifiedUser.UsersAskedFeed, feed)

	//The last step is add the feedback to the slice.
	// isFeedInserted, err := db.ModifyUser(modifiedUser, rID)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Cannot modify the feedback status on this user")
	// 	return
	// }

	// fmt.Println(isFeedInserted)

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
	if count == len(group) {
		return false
	}
	return true
}

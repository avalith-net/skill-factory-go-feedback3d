package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/gin-gonic/gin"

	services "github.com/avalith-net/skill-factory-go-feedback3d/services"
)

const (
	timeFormat = "2006.01.02 15:04:05"
)

// RequestFeedback godoc
// @Description get string by id
// @id RequestFeedback
// @Summary is used to request a feedback to other user.
// @Param id query string true "Account ID"
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /fbRequest [post]
func RequestFeedback(c *gin.Context) {
	id := c.Query("id")
	if len(id) < 1 || id == IDUser {
		c.String(http.StatusBadRequest, "Error with the request.")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User not found.")
		return
	}

	loggedUser, err := db.GetUser(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exists.")
		return
	}
	//------------------------------------------------------------------------------------------

	if !user.Enabled {
		c.String(http.StatusBadRequest, "Cannot ask feedback, the target user is banned")
		return
	}

	//We searching for a feedback request between the logged user and requested user.
	// In case it does exist, we calculate the diff of seconds and days.
	// If 15 days has passed, the user can request another feedback.
	// Also if it is the first request between this 2 users, request its allow, that's
	//why we check for the seconds. If the diff is 0, it means there is no previous request.
	reqID, found := db.GetFeedBackRequestedID(id, IDUser)
	reqObj, _, _ := db.GetSelectedFeedBackRequestObj(reqID)

	if found {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You already have a request pending from this user", "timeleft": reqObj.TimeLeft})
		return
	}

	var feedRequested models.FeedbacksRequested

	feedRequested.RequestedUserID = id
	feedRequested.UserLoggedID = IDUser
	feedRequested.RequestedUserName = user.Name
	feedRequested.RequestedUserLastName = user.LastName
	feedRequested.SentDate = time.Now()
	feedRequested.TimeLeft = 15

	_, isRequestCreated, err := db.AddFeedbackRequested(feedRequested)
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot create a feedback request to target user in feedback request.")
		return
	}
	if !isRequestCreated {
		c.String(http.StatusBadRequest, "Something went wrong trying create a request to target user in feedback request.")
		return
	}

	var userAskingFeed models.UsersAskedFeed

	userAskingFeed.UserWhoAskFeedID = IDUser
	userAskingFeed.UserAskedForFeedID = id
	userAskingFeed.NameWhoAskFeed = loggedUser.Name
	userAskingFeed.LastNameWhoAskFeed = loggedUser.LastName
	userAskingFeed.SentDate = time.Now()

	userAskingFeedID, isRequestCreated, err := db.AddUsersAsksFeed(userAskingFeed)
	if err != nil {
		c.String(http.StatusBadRequest, "Cannot create a userAsksFeed instance in feedback request.")
		return
	}
	if !isRequestCreated {
		c.String(http.StatusBadRequest, "Something went wrong trying create a userAsksFeed instance in feedback request. ID: "+userAskingFeedID)
		return
	}

	//------------------------------------------------------------------------------------------
	// An email is send to all admins with the report notification.
	bodyString := "Hi <b><i>" + user.Name + "</i></b>!\n" +
		"I'd like to ask a few questions about your working experience with me. It's important to help me to improve." +
		"Follow this link to give me feedback: <b><i>http:localhost:8080/feedback?target_id=" + IDUser +
		"</i></b>\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format(timeFormat)

	//Email send function
	go func() {
		if !services.SendEmail(user.Email, "Feedback request.", bodyString) {
			log.Println("email could not be sent")
		}
	}()

	c.String(http.StatusCreated, "Success")
}

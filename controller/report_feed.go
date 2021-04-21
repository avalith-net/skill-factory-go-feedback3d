package controller

import (
	"net/http"
	"time"

	services "github.com/avalith-net/skill-factory-go-feedback3d/services"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

// Report feedback godoc
// @Description it lets you report a feedback
// @id reportfb
// @Summary is used to report a feedback to an admin. as a user
// @Param Authorization header string true "jwt token"
// @Param feed_id path string true "id"
// @Success 201 {string} string "Feed reported"
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Unauthorized"
// @Failure 500 {string} string "An error has ocurred trying to report feedback."
// @Failure default {string} string "An error has ocurred"
// @Router /users/report/{feed_id} [patch]
func ReportFeed(c *gin.Context) {
	feedID := c.Param("feed_id")

	feed, err := db.GetSelectedFeedBack(feedID)
	if err != nil {
		c.String(http.StatusBadRequest, "Feed does not exist.")
		return
	}

	if feed.IsApprobed {
		c.String(http.StatusBadRequest, "You're trying to access an approbed feed.")
		return
	}

	if feed.IsReported {
		c.String(http.StatusBadRequest, "This feedback was reported already. Can not report twice the same feedback.")
		return
	}

	isReported, err := db.ReportFeedback(feed, feedID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error trying to report feedback.")
		return
	}

	if !isReported {
		c.String(http.StatusBadRequest, "Feedback could not be reported.")
		return
	}
	admins, err := db.GetAllAdmins()
	_, err = db.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Coudn´t get the admins": err.Error()})
		return
	}

	// We go through each administrator to access each token and each email to be able to send the reported feeback
	for _, eachAdmin := range admins {

		bodyString := "<b><i>Hi admin!</i></b>\n" +
			"A user ask yours review about a feedback. Could you please take a moment to check it out?" +
			"Follow this link to give me feedback: <b><i>http:localhost:8080/selectedFeedback?feedID=" + feedID +
			"</i></b>\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format("2006.01.02 15:04:05")

		if !services.SendEmail(eachAdmin.Email, "Feedback reported.", bodyString) {
			c.String(http.StatusBadRequest, "An error has ocurred sending the email "+err.Error())
			return
		}
	}

	c.String(http.StatusCreated, "Feedback reported.")
}

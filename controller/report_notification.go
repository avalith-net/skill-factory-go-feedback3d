package controller

import (
	"net/http"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	services "github.com/JoaoPaulo87/skill-factory-go-feedback3d/services/email"
	"github.com/gin-gonic/gin"
)

//ReportNotification godoc
// @Summary Report feedback
// @Description Get all the users with admin. rol from the db and send them a feedback report
// @User get-struct-by-json
// @Accept  json
// @Param feedID query string true "insert feedback ID here"
// @Param Authorization header string true "Token"
// @Success 200 {string} string "status ok"
// @Failure 400 {string} string "Page must be a number"
// @Failure 400 {string} string "Coudn´t send report"
// @Failure default {string} string "Error"
// @Router /reportNotification [post]
func ReportNotification(c *gin.Context) {
	//At the moment, report_feed is doing the same thing. Maybe we could delete this.
	admins, err := db.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Coudn´t get the admins": err.Error()})
		return
	}

	FeedID := c.Query("feedID")
	if err != nil {
		c.String(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	// We go through each administrator to access each token and each email to be able to send the reported feeback
	for _, eachAdmin := range admins {

		bodyString := "<b><i>Hi admin!</i></b>\n" +
			"A user ask yours review about a feedback. Could you please take a moment to check it out?" +
			"Follow this link to give me feedback: <b><i>http:localhost:8080/selectedFeedback?feedID=" + FeedID +
			"</i></b>\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format("2006.01.02 15:04:05")

		if !services.SendEmail(eachAdmin.Email, "Feedback reported.", bodyString) {
			c.String(http.StatusBadRequest, "An error has ocurred sending the email "+err.Error())
			return
		}
	}

	c.String(http.StatusCreated, "Mail sended!")
}

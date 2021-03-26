package controller

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"

	services "github.com/blotin1993/feedback-api/services/email"
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

	bodyString := "Hi <b><i>" + user.Name + "</i></b>!\n" +
		"I'd like to ask a few questions about your working experience with me. It's important to help me to improve." +
		"Follow this link to give me feedback: <b><i>http:localhost:8080/feedback?target_id=" + IDUser +
		"</i></b>\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format("2006.01.02 15:04:05")

	//Email send function
	if !services.SendEmail(user.Email, "Feedback request.", bodyString) {
		c.String(http.StatusBadRequest, "An error has ocurred sending the email")
		return
	}
	c.String(http.StatusCreated, "Success")
}

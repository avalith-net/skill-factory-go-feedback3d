package routers

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"

	services "github.com/blotin1993/feedback-api/services/email"
)

// RequestFeedback godoc
// @Description get string by id
// @id changepass
// @Summary is used to request a feedback to other user.
// @Param id query string true "Account ID"
// @Produce plain
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Wrong mail error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /fbRequest [post]
func RequestFeedback(c *gin.Context) {
	var u models.FeedbackRequest
	var userR models.ReturnUser

	u.ReceiverID = c.Query("id")

	userR, err := db.GetUser(u.ReceiverID)
	if e != nil {
		http.Error(w, "internal error!", http.StatusBadRequest)
		return
	}

	u.ReceiverEmail = userR.Email

	if len(u.ReceiverEmail) < 1 {
		c.String(http.StatusBadRequest, "must complete email form")
		return
	}

	auxUser, mailExist, _ := db.UserAlreadyExist(u.ReceiverEmail)
	if !mailExist {
		c.String(http.StatusBadRequest, "Wrong mail.")
		return
	}

	bodyString := "Hi <b><i>" + auxUser.Name + "</i></b>!\n" +
		"I'd like to ask a few questions about your working experience with me. It's important to help me to improve.\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format("2006.01.02 15:04:05")

	//Email send function
	if !services.SendEmail(auxUser.Email, "Feedback request.", bodyString) {
		c.String(http.StatusBadRequest, "An error has ocurred sending the email")
		return
	}
	c.String(http.StatusCreated, "Success")
}

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
// @id RequestFeedback
// @Summary is used to request a feedback to other user.
// @Param id query string true "Account ID"
// @Header token string true "Token"
// @Produce plain
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /fbRequest [post]
func RequestFeedback(c *gin.Context) {
	var userID string
	var userR models.ReturnUser

	userID := c.Query("id")

	userR, err := db.GetUser(userID)
	if err != nil {
		http.Error(w, "internal error!", http.StatusBadRequest)
		return
	}

	bodyString := "Hi <b><i>" + userR.Name + "</i></b>!\n" +
		"I'd like to ask a few questions about your working experience with me. It's important to help me to improve.\n<br> Thanks for your time!\n\n<br><b> Feedback-Api</b> \n <br><i>feedbackapiadm@gmail.com</i>\n<br> " + time.Now().Format("2006.01.02 15:04:05")

	//Email send function
	if !services.SendEmail(userR.Email, "Feedback request.", bodyString) {
		c.String(http.StatusBadRequest, "An error has ocurred sending the email")
		return
	}
	c.String(http.StatusCreated, "Success")
}

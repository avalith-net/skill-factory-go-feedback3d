package routers

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"

	services "github.com/blotin1993/feedback-api/services/email"
)

//RequestFeedback func
func RequestFeedback(w http.ResponseWriter, r *http.Request) {
	var u models.FeedbackRequest
	var user models.ReturnUser

	u.ReceiverID = r.URL.Query().Get("id")

	user, e := db.GetUser(u.ReceiverID)
	if e != nil {

	}

	u.ReceiverEmail = user.Email

	if len(u.ReceiverEmail) < 1 {
		http.Error(w, "must complete email form", http.StatusBadRequest)
		return
	}

	auxUser, mailExist, _ := db.UserAlreadyExist(u.ReceiverEmail)
	if !mailExist {
		http.Error(w, "Wrong mail.", 400)
		return
	}

	bodyString := "Hi <b><i>" + auxUser.Name + "</i></b>!\n" +
		"I'd like to ask a few questions about your working experience with me. It's important to help me to improve.\n Thanks for your time!\n\n Feedback-Api \n <i>feedbackapiadm@gmail.com</i>\n " + time.Now().Format("2006.01.02 15:04:05")

	//Email send function
	var err error
	if !services.SendEmail(auxUser.Email, "Feedback request.", bodyString) {
		http.Error(w, "An error has ocurred sending the email"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

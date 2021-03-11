package routers

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	jwt "github.com/blotin1993/feedback-api/services/auth"
	services "github.com/blotin1993/feedback-api/services/email"
)

//RecoverPass - receive the user data from DataBase and send an Email with his current password
func RecoverPass(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if len(email) < 1 {
		http.Error(w, "must complete email form", http.StatusBadRequest)
		return
	}
	user, mailExist, _ := db.UserAlreadyExist(email)
	if !mailExist {
		http.Error(w, "Wrong mail.", 400)
		return
	}

	stringObjectID := (user.ID).Hex()
	expirationTime := time.Now().Add(1 * time.Hour)
	jwtKey, err := jwt.GenerateJWT(user, expirationTime)

	bodyString := "Hey <b><i>" + user.Name + "</i></b>!\nFollow this link to recover your password.\n <a>http:localhost/8080/changePassword?id=" + stringObjectID + "&token=Bearer " + jwtKey + "</a>"

	//Email send function
	if !services.SendEmail(email, "Get your password.", bodyString) {
		http.Error(w, "An error has ocurred sending the email"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

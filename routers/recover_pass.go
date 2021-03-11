package routers

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	services "github.com/blotin1993/feedback-api/services/email"
)

//GetMyPassword - receive the user data from DataBase and send an Email with his current password
func GetMyPassword(w http.ResponseWriter, r *http.Request) {
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
	decryptedPassword, err := db.DecryptPassword(user.Password)
	if err != nil {
		http.Error(w, "Internal error"+err.Error(), 400)
		return
	}
	bodyString := "Hola <b><i>" + user.Name + "</i></b>!\n" +
		"Recibimos tu pedido de recuperacion de password! \n Your Password: " + decryptedPassword +
		" \n\n Esperamos que sea de tu ayuda, estamos a tu disposicion! \n\n Feedback-Api Admin\n <i>feedbackapiadm@gmail.com</i>\n " +
		time.Now().Format("2006.01.02 15:04:05")

	//Email send function
	if !services.SendEmail(email, "Get your password.", bodyString) {
		http.Error(w, "An error has ocurred sending the email"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

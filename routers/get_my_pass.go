package routers

import (
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"gopkg.in/gomail.v2"
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
	msg := gomail.NewMessage()
	msg.SetHeader("From", "vlotingaming@gmail.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Get your password.")
	msg.SetBody("text/html", "Hola <b><i>"+user.Name+"</i></b>!\n"+
		"Recibimos tu pedido de recuperacion de password! \n Your Password: "+decryptedPassword+
		" \n\n Esperamos que sea de tu ayuda, estamos a tu disposicion! \n\n Feedback-Api Admin\n <i>feedbackapiadm@gmail.com</i>\n "+
		time.Now().Format("2006.01.02 15:04:05"))

	//Send the email to user
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "vlotingaming@gmail.com", "V2H!x%CaxCeM")
	if err := d.DialAndSend(msg); err != nil {
		http.Error(w, "An error has ocurred sending the email"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

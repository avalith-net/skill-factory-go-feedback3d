package routers

import (
	"net/http"
	"time"

	"gopkg.in/gomail.v2"
)

//GetMyPassword - receive the user data from DataBase and send an Email with his current password
func GetMyPassword(w http.ResponseWriter, r *http.Request) {
	/* ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	@@ -22,15 +22,27 @@ func GetMyPassword(w http.ResponseWriter, r *http.Request) {
		return
	}
	decryptedPassword, err := db.DecryptPassword(profile.Password)
	if err != nil {
		http.Error(w, "Ocurrio un error al desencriptar la contrasena"+err.Error(), 400)
		return
	} */

	msg := gomail.NewMessage()
	msg.SetHeader("From", "feedbackapiadm@gmail.com")
	msg.SetHeader("To", "abotlucasmdq@gmail.com")
	msg.SetHeader("Subject", "Get your password.")
	msg.SetBody("text/html", "Hola <b><i>Lucas</i></b>/n"+
		"Recibimos tu pedido de recuperacion de password! \n Your Password: 123456" /* +decryptedPassword+ */ +
		" \n\n Esperamos que sea de tu ayuda, estamos a tu disposicion! \n\n Feedback-Api Admin\n <i>feedbackapiadm@gmail.com</i>\n "+
		time.Now().Format("2006.01.02 15:04:05"))

	//Send the email to user
	d := gomail.NewPlainDialer("smtp.gmail.com", 465, "feedbackapiadmi1@gmail.com", "feedback123")
	if err := d.DialAndSend(msg); err != nil {
		http.Error(w, "Ocurrio un error al enviar el mail"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

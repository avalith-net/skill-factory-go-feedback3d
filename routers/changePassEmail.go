package routers

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
)

//ChangePassEmail .
func ChangePassEmail(w http.ResponseWriter, r *http.Request) {
	//parámetros: id, token, newpass
	id := r.URL.Query().Get("id")
	token := r.URL.Query().Get("token")
	newPass := r.URL.Query().Get("newpass")
	if len(newPass) < 6 {
		http.Error(w, "The new password must be at least 6 characters long", 400)
		return
	}
	user := models.User{
		Password: newPass,
	}
	_, isOk, _, _ := TokenProcess(token)
	if !isOk {
		http.Error(w, "Authentication error.", 400)
		return
	}
	//modificar usuario
	hasEffect, err := db.ModifyUser(user, id)
	if !hasEffect {
		http.Error(w, "An error has ocurred trying to set a new password."+err.Error(), 400)
	}
	w.WriteHeader(http.StatusCreated)
}

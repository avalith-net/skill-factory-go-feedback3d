package routers

import (
	"encoding/json"
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
)

//SignUp is used to register to the app
func SignUp(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error checking the data: "+err.Error(), 400)
		return
	}

	//data validation
	if len(t.Email) == 0 {
		http.Error(w, "must add email ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Your password must be at least 6 characters long", 400)
		return
	}
	_, encontrado, _ := db.UserAlreadyExist(t.Email)
	if encontrado == true {
		http.Error(w, "Email already registered.", 400)
		return
	}
	_, status, err := db.AddRegister(t)
	if err != nil {
		http.Error(w, "Database error "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Error, Register not added.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	jwt "github.com/blotin1993/feedback-api/services/auth"
)

//Login validation
func Login(w http.ResponseWriter, r *http.Request) {

	var usu models.User
	err := json.NewDecoder(r.Body).Decode(&usu)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//email validation
	if len(usu.Email) == 0 {
		http.Error(w, "Email needed.", 400)
		return
	}
	document, exists := db.LoginAttempt(usu.Email, usu.Password)
	if exists == false {
		http.Error(w, "Wrong user or password.", 400)
		return
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	jwtKey, err := jwt.GenerateJWT(document, expirationTime)
	if err != nil {
		http.Error(w, "Error generating the token "+err.Error(), 400)
		return
	}

	resp := models.LoginReply{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//cookie set for expirationTime var time
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

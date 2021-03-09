package db

import (
	"github.com/blotin1993/feedback-api/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginAttempt checks if the user already exists and verifies the password.
func LoginAttempt(email string, password string) (models.User, bool) {
	user, found, _ := UserAlreadyExist(email)
	if found == false {
		return user, false
	}

	passwordToBytes := []byte(password)                               // the param pass
	passwordDb := []byte(user.Password)                               // db pass
	err := bcrypt.CompareHashAndPassword(passwordDb, passwordToBytes) // comparing both

	if err != nil {
		return user, false
	}
	return user, true
}

package db

import (
	"github.com/avalith-net/skill-factory-go-feedback3d/auth"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
)

//LoginAttempt checks if the user already exists and verifies the password.
func LoginAttempt(email string, password string) (models.User, bool) {
	user, found, _ := UserAlreadyExist(email)
	if !found {
		return user, false
	}

	passwordToBytes := []byte(password) // the param pass
	passwordDb := []byte(user.Password) // db pass
	match, _ := auth.ComparePasswords(passwordToBytes, passwordDb)

	if !match {
		return user, false
	}
	return user, true
}

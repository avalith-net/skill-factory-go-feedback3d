package db

import (
	"github.com/blotin1993/feedback-api/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin tries to connect the API with the database*/
func IntentoLogin(email string, password string) (models.User, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	// Now i compares if the password received matches with the password in the database
	// i create a variable slice of bytes
	passwordBytes := []byte(password)
	// i create another variable with the password i got in the database
	passwordBD := []byte(usu.Password)
	// Ahora llamo a una función del package bcrypt que compara las password
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}

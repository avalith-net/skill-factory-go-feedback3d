package db

import "golang.org/x/crypto/bcrypt"

//PassEncrypt is used to encrypt the user password before it goes to de db.
func PassEncrypt(pass string) (string, error) {
	cost := 8 // 2 pow cost to generate our pass hash.
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

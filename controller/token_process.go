package controller

import (
	"errors"
	"os"
	"strings"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email is gonna be used to proccess the token
var Email string

//IDUser will be given by the jwt proccessing
var IDUser string

//TokenProcess will validate our token
func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	claims := &models.Claim{} // jwt structure

	splitToken := strings.Split(tk, "Bearer") // we need  to remove the "Bearer" from the token

	if len(splitToken) != 2 { // [Bearer],[(token)] else -> error
		return claims, false, string(""), errors.New("invalid token format")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	//---> valid token --> global variables set up.
	if err == nil {
		_, found, _ := db.UserAlreadyExist(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}

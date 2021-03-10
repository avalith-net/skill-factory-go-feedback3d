package jwt

import (
	"os"
	"time"

	"github.com/blotin1993/feedback-api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera un jwt para loguear.
func GeneroJWT(usu models.User) (string, error) {
	miClave := []byte(os.Getenv("JWT_KEY"))

	payload := jwt.MapClaims{
		"email":    usu.Email,
		"name":     usu.Name,
		"lastname": usu.LastName,
		"_id":      usu.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

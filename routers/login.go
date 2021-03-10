package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/jwt"
	"github.com/blotin1993/feedback-api/models"
)

/*Login realiza el login de usuario
Recibe como parámetro lo mismo de tods los endpoints y no devuelve nada,
como los otros endPoints, son prácticamente métodos*/
func Login(w http.ResponseWriter, r *http.Request) {
	// Vamos a setear en el header que el contenido que devolveremos (w)
	// será de tipo Json
	w.Header().Add("content-type", "application/json")

	var usu models.User
	err := json.NewDecoder(r.Body).Decode(&usu)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(usu.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	documento, existe := db.IntentoLogin(usu.Email, usu.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}
	// Si el token está generado:
	resp := models.LoginReply{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Vamos también a grabar una cookie
	// generamos un campo fecha para ver la expiración de esa cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

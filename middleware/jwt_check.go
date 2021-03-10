package middleware

import (
	"net/http"

	"github.com/blotin1993/feedback-api/routers"
)

//ValidoJWT ..
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

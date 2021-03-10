package middleware

import (
	"net/http"

	"github.com/blotin1993/feedback-api/routers"
)

//ValidateJWT is used to check the jwt passed as parameter.
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.TokenProcess(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token error."+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

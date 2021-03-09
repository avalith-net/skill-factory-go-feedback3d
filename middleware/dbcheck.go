package middleware

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
)

//CheckDb middleware.
func CheckDb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection lost.", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

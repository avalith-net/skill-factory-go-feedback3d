package routers

import (
	"encoding/json"
	"net/http"

	"github.com/blotin1993/feedback-api/db"
)

//GetFeed is used to set the profile picture
func GetFeed(w http.ResponseWriter, r *http.Request) {
	feedSlice, err := db.GetFeedFromDb(IDUser)
	if err != nil {
		http.Error(w, "Error"+err.Error(), 400)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(feedSlice)
}

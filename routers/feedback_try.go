package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
)

//FeedbackTry is used to process our feedbacks
func FeedbackTry(w http.ResponseWriter, r *http.Request) {
	rID := r.URL.Query().Get("target_id")
	if len(rID) < 1 {
		http.Error(w, "ID Error", http.StatusBadRequest)
		return
	}
	var fb models.Feedback

	err := json.NewDecoder(r.Body).Decode(&fb)

	//-------Feedback validation------
	// n := structs.Map(fb.PerformanceArea)
	// fmt.Println(n)
	// httpServer := fb.Field("Server").Fields()

	// for x, v := range n {
	// 	fmt.Println(x, v)
	// }

	//-----------------------------------

	fb.IssuerID = IDUser
	fb.ReceiverID = rID
	fb.Date = time.Now()

	_, status, err := db.AddFeedback(fb)

	//si hay un error
	if err != nil {
		http.Error(w, "An error has ocurred. Try again later "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Database error.", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

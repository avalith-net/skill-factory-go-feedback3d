package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"

	"github.com/fatih/structs"
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
	if structs.IsZero(fb) || !hasZeroGroup(fb.PerformanceArea, fb.TeamArea, fb.TechArea) {
		http.Error(w, "You must enter at least one complete area", 400)
		return
	}

	if !validateMsgLength(1614, fb.Message) {
		http.Error(w, "Message cannot be longer than 1500 characters.", 400)
		return
	}
	if !validateMsgLength(536, fb.TechArea.Message, fb.TeamArea.Message, fb.PerformanceArea.Message) { //36 porque toma 12 más por salto de página (en este caso serían 3, checkear. No estoy seguro si es solo en postman)
		http.Error(w, "Area Messages cannot be longer than 500 characters.", 400)
		return
	}
	//-----------------------------------

	fbProcessed := &models.Feedback{}

	fbProcessed.IssuerID = IDUser
	fbProcessed.ReceiverID = rID
	fbProcessed.Date = time.Now()

	_, status, err := db.AddFeedback(*fbProcessed)

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

func validateMsgLength(maxLen int, Amsg ...string) bool {
	for _, msg := range Amsg {
		if len(msg) > maxLen {
			return false
		}
	}
	return true
}

func hasZeroGroup(gr ...interface{}) bool {
	count := 0
	for _, field := range gr {
		if structs.HasZero(field) {
			count++
		}
	}
	if count == len(gr) {
		return false
	}
	return true
}

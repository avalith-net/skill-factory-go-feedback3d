package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/ulule/deepcopier"

	"github.com/fatih/structs"
)

//FeedbackTry is used to process our feedbacks
func FeedbackTry(w http.ResponseWriter, r *http.Request) {
	rID := r.URL.Query().Get("target_id")
	if len(rID) < 1 {
		http.Error(w, "ID Error", http.StatusBadRequest)
		return
	}
	var fbRaw models.FeedbackRaw

	err := json.NewDecoder(r.Body).Decode(&fbRaw)

	//-------Feedback validation------
	if structs.IsZero(fbRaw) ||
		(structs.IsZero(fbRaw.PerformanceArea) && structs.IsZero(fbRaw.TeamArea) && structs.IsZero(fbRaw.TechArea)) ||
		(structs.HasZero(fbRaw.PerformanceArea) && structs.HasZero(fbRaw.TeamArea) && structs.HasZero(fbRaw.TechArea)) {
		http.Error(w, "You must enter at least one complete area", 400)
		return
	}
	//-----------------------------------

	fbProcessed := &models.Feedback{}
	deepcopier.Copy(fbRaw).To(fbProcessed)

	fbProcessed.IssuerID = IDUser
	fbProcessed.ReceiverID = rID
	fbProcessed.Date = time.Now()

	_, status, err := db.AddFeedback(*fbProcessed)

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

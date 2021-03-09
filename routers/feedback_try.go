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
	//
	fb.IssuerID = IDUser
	fb.ReceiverID = rID
	fb.Date = time.Now()

	//-------Feedback validation------<<<<<<<<<<<<<<<<<<<<<<<<<<<<< TO IMPROVE
	if len(fb.TeamArea.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}
	if len(fb.TeamArea.Message) >= 500 {
		http.Error(w, "must have a maximum of 500 characters", 400)
		return
	}

	if len(fb.TeamArea.TeamPlayer) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TeamArea.Commited) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TeamArea.Communication) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	//validation TechArea
	if len(fb.TechArea.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}

	if len(fb.TechArea.Message) >= 500 {
		http.Error(w, "must have a maximum of 500 characters", 400)
		return
	}

	if len(fb.TechArea.BestPractices) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TechArea.TechKnowledge) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.TechArea.CodingStyle) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	//validation PerformanceArea
	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}

	if len(fb.PerformanceArea.Message) >= 500 {
		http.Error(w, "must have a maximum of 500 characters", 400)
		return
	}

	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	if len(fb.PerformanceArea.Message) == 0 {
		http.Error(w, "empty field", 400)
		return
	}

	//Validation Message
	if len(fb.Message) == 0 {
		http.Error(w, "the message must have at least one character", 400)
		return
	}

	if len(fb.Message) >= 1500 {
		http.Error(w, "must have a maximum of 1500 characters", 400)
		return
	}
	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

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

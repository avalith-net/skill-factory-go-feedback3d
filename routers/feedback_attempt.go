package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"

	"github.com/fatih/structs"
)

//FeedbackTry is used to process our feedbacks
func FeedbackTry(c *gin.Context) {
	rID := c.Query("target_id")
	if len(rID) < 1 {
		c.String(http.StatusBadRequest, "ID Error")
		return
	}
	// user, err := db.GetUser(rID)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "User "+user.Name+"Not registered.")
	// 	return
	// }
	var fb models.Feedback
	err := json.NewDecoder(c.Request.Body).Decode(&fb)
	if err != nil {
		c.String(http.StatusBadRequest, "User Not registered.")
		return
	}

	//-------Feedback validation------
	if structs.IsZero(fb) || !hasZeroGroup(fb.PerformanceArea, fb.TeamArea, fb.TechArea) {
		c.String(http.StatusBadRequest, "You must enter at least one complete area")
		return
	}

	if !validateMsgLength(1615, fb.Message) { //1615 xq toma saltos de pagina como caracteres.
		c.String(http.StatusBadRequest, "Message cannot be longer than 1500 characters")
		return
	}
	if !validateMsgLength(540, fb.TechArea.Message, fb.TeamArea.Message, fb.PerformanceArea.Message) { //40 xq toma saltos de página
		c.String(http.StatusBadRequest, "Area Messages cannot be longer than 500 characters.")
		return
	}
	//-----------------------------------

	fb.IssuerID = IDUser
	fb.ReceiverID = rID
	fb.Date = time.Now()

	_, status, err := db.AddFeedback(fb)

	if err != nil {
		c.String(http.StatusInternalServerError, "An error has ocurred. Try again later "+err.Error())
		// c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if status == false {
		c.String(http.StatusInternalServerError, "Database error.")
		return
	}
	c.String(http.StatusCreated, "Success")
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

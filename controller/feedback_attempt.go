package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"

	"github.com/fatih/structs"
)

// Feedback Attempt godoc
// @Description give feedback
// @id feedattempt
// @Summary is used to give a feedback to a user.
// @Param id query string true "target ID"
// @Param token query string true "JWT Token"
// @Produce plain
// @Success 201 {string} string "Email sended successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /fbRequest [post]
func FeedbackAttempt(c *gin.Context) {
	rID := c.Query("target_id")
	if len(rID) < 1 {
		c.String(http.StatusBadRequest, "ID Error")
		return
	}
	user, err := db.GetUser(rID)
	_, isFound, _ := db.UserAlreadyExist(user.Email)
	if !isFound {
		c.String(http.StatusBadRequest, "User was not found.")
		return
	}
	var fb models.Feedback
	err = json.NewDecoder(c.Request.Body).Decode(&fb)
	if err != nil {
		c.String(http.StatusBadRequest, "Check your form.")
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
	if !validateMsgLength(541, fb.TechArea.Message, fb.TeamArea.Message, fb.PerformanceArea.Message) { //40 xq toma saltos de página
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

func hasZeroGroup(group ...interface{}) bool {
	count := 0
	for _, field := range group {
		if !structs.HasZero(field) {
			return true
		}
		count++
	}
	if count == len(group) {
		return false
	}
	return true
}

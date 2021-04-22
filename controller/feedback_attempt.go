package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/avalith-net/skill-factory-go-feedback3d/services"
	"github.com/gin-gonic/gin"

	"github.com/fatih/structs"
)

// Feedback  godoc
// @Description gives feedback to the user
// @id fb
// @Summary is used to give feedback to users.
// @Param target_id query string true "Target ID"
// @Param feedback body string true "Json body with email and password"
// @Param Authorization header string true "JWT Token"
// @Accept  json
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Wrong mail or password."
// @Failure 500 {string} string "Error generating the token."
// @Failure default {string} string "An error has ocurred"
// @Router /feedback [post]
func FeedbackTry(c *gin.Context) {

	validUser, _ := db.GetUser(IDUser)
	if validUser.Enabled == false {
		c.String(http.StatusUnauthorized, "User not authorized.")
		return
	}

	rID := c.Query("target_id")
	if len(rID) < 1 {
		c.String(http.StatusBadRequest, "ID Error")
		return
	}
	user, err := db.GetUser(rID)
	_, isFound, _ := db.GetUserByEmail(user.Email)
	if !isFound {
		c.String(http.StatusBadRequest, "User was not found.")
		return
	}
	if user.Enabled == false {
		c.String(http.StatusUnauthorized, "User not authorized to receive feedbacks.")
		return
	}

	var fb models.Feedback

	if err := c.ShouldBindJSON(&fb); err != nil {
		c.String(http.StatusBadRequest, "Check your form.")
		return
	}

	//-------Feedback validation------
	if structs.IsZero(fb) || !hasZeroGroup(fb.PerformanceArea, fb.TeamArea, fb.TechArea) {
		c.String(http.StatusBadRequest, "You must enter at least one complete area")
		return
	}
	//-----------------------------------Format validation

	isValid, err := services.ValidateForm(fb)
	if !isValid {
		c.String(http.StatusBadRequest, "Invalid format: "+err.Error())
		return
	}

	// graphic stats
	err = services.InitGraphic(fb, &user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(user.Graphic)
	// persist graphic.
	_, err = db.UpdateGraphic(user, rID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------

	fb.IssuerID = IDUser
	fb.IssuerName = validUser.Name + " " + validUser.LastName
	fb.ReceiverID = rID
	fb.Date = time.Now()

	// Send email notification

	msg := "Hi " + user.Name + " " + user.LastName + "! \n You have received a new feedback, check it in your dashboard! <a>http:localhost:8080/dashboard</a> \n\n"
	services.SendEmail(user.Email, "New Feedback Received!", msg)

	//------------------------------

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

func hasZeroGroup(group ...interface{}) bool {
	count := 0
	for _, field := range group {
		if !structs.HasZero(field) {
			return true
		}
		count++
	}
	return count != len(group)
}

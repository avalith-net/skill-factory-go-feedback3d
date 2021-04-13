package controller

// seria un post. Si esta aprobado, cambia el estado del feedback isApprobed a true, isReported a false y isDisabled a false. En caso contrario, todo lo opuesto.

import (
	"net/http"
	"strconv"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

// FeedbackState godoc
// @Description it creates a feedback state
// @id feedState
// @Summary Used when an administrator evaluates reported feedback
// @Param Authorization header string true "jwt token"
// @Param is_approbed path bool true "is_approbed"
// @Param feed_id query string true "insert feedback ID here"
// @Success 201 {string} string "Feed status created"
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Unauthorized"
// @Failure 500 {string} string "An error has ocurred trying to set feedback state."
// @Failure default {string} string "An error has ocurred"
// @Router /users/feedState/{is_approbed} [patch]
func FeedbackState(c *gin.Context) {
	//This endpoint logic it's when an admin. recives a reported feed, a page with the feed and 2 buttons appears,
	//a button to approbe and another to disapprobe. When that happens it redirected to this endpoint with the
	//result and if the feed is approbed, the report turns "false" and cant be reported again once approbed.
	//But if its disapprobed, then the feedback is delete from the db.
	feedState := c.Param("is_approbed")
	feedID := c.Query("feed_id")

	feed, err := db.GetSelectedFeedBack(feedID)
	if err != nil {
		c.String(http.StatusBadRequest, "Feed does not exist.")
		return
	}

	if feed.IsApprobed {
		c.String(http.StatusBadRequest, "You're trying to access an approbed feed.")
		return
	}

	if !feed.IsReported {
		c.String(http.StatusBadRequest, "You're trying to access to an unreported feedback.")
		return
	}

	state, _ := strconv.ParseBool(feedState)

	isApprobedState, err := db.UpdateFeedbackState(feed, feedID, state)
	if err != nil || !isApprobedState {
		c.String(http.StatusBadRequest, "Database error "+err.Error())
		return
	}

	if state {
		c.String(http.StatusCreated, "Feedback state created. Feed approbed")
	} else {
		deleteResponse, _ := db.DeleteFeedback(feedID)

		c.JSON(http.StatusCreated, gin.H{"Feedback state created. Feed was disapprobed and deleted. ": deleteResponse})
	}
}

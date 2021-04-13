package controller

import (
	"net/http"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

// Report feedback godoc
// @Description it lets you report a feedback
// @id reportfb
// @Summary is used to report a feedback to an admin. as a user
// @Param Authorization header string true "jwt token"
// @Param feed_id path string true "id"
// @Success 201 {string} string "Feed reported"
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Unauthorized"
// @Failure 500 {string} string "An error has ocurred trying to report feedback."
// @Failure default {string} string "An error has ocurred"
// @Router /users/report/{feed_id} [patch]
func ReportFeed(c *gin.Context) {
	feedID := c.Param("feed_id")

	feed, err := db.GetSelectedFeedBack(feedID)
	if err != nil {
		c.String(http.StatusBadRequest, "Feed does not exist.")
		return
	}

	if feed.IsApprobed {
		c.String(http.StatusBadRequest, "You're trying to access an approbed feed.")
		return
	}

	isReported, err := db.ReportFeedback(feed, feedID)
	if err != nil || !isReported {
		c.String(http.StatusInternalServerError, "Error trying to report feedback")
		return
	}

	c.String(http.StatusCreated, "Feedback reported.")
}

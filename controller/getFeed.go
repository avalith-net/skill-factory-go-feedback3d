package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

// GetFeed godoc
// @Description get all feedbacks from db.
// @id getfeed
// @Summary is used to get al the feeds of a user.
// @Param id query string true "target ID"
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Success"
// @Failure 400 {string} string "Error ID."
// @Failure default {string} string "An error has ocurred"
// @Router /getfb [get]
func GetFeed(c *gin.Context) {
	feedSlice, err := db.GetFeedFromDb(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "Error"+err.Error())
		return
	}
	c.JSON(http.StatusCreated, feedSlice)
}

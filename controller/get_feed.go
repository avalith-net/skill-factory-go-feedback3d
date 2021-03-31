package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

// dashboard godoc
// @Description get user feedbacks from db.
// @id dash
// @Summary is used to get feedbacks of the user.
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Success"
// @Failure 400 {string} string "Error ID."
// @Failure default {string} string "An error has ocurred"
// @Router /dashboard [get]
func GetFeed(c *gin.Context) {
	feedSlice, err := db.GetFeedFromDb(IDUser, true)
	if err != nil {
		c.String(http.StatusBadRequest, "Error"+err.Error())
		return
	}
	c.JSON(http.StatusCreated, feedSlice)
}

package controller

import (
	"net/http"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

// selectedFeedback godoc
// @Description get one user fb from db.
// @id selectedFeedback
// @Summary is used to get one selected feedback of the user.
// @Param Authorization header string true "jwt token"
// @Param feedID query string true "insert feedback ID here"
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Success"
// @Failure 400 {string} string "Error ID."
// @Failure default {string} string "An error has ocurred"
// @Router /selectedFeedback [get]
func GetSelectedFeedback(c *gin.Context) {
	FeedID := c.Query("feedID")
	feed, err := db.GetSelectedFeedBack(FeedID)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	c.JSON(http.StatusCreated, feed)
}

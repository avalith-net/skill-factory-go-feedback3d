package controller

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

// getAllMyFeedbacks godoc
// @Description get all my feedbacks
// @id GetAllMyFeedbacks
// @Summary is used to get all the feedbacks the users send me.
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Feedbacks returned successfully."
// @Header 201 {string} string "Success"
// @Failure 400 {string} string "Error ID."
// @Failure default {string} string "An error has ocurred"
// @Router /users/my_feedbacks [get]
func GetAllMyFeedbacks(c *gin.Context) {
	allMyFeeds, err := db.GetFeedFromDb(IDUser, true)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	c.JSON(http.StatusCreated, allMyFeeds)
}

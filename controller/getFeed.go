package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

//GetFeed is used to set the profile picture
func GetFeed(c *gin.Context) {
	feedSlice, err := db.GetFeedFromDb(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "Error"+err.Error())
	}
	c.JSON(http.StatusCreated, feedSlice)
}

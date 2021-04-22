package controller

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

// getFeedRequestsFromMe godoc
// @Description get all feedbacks requests from me
// @id getFeedRequestsFromMe
// @Summary is used to get all the feedbacks requests the users send me.
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Requests returned successfully."
// @Header 201 {string} string "Success"
// @Failure 400 {string} string "Error ID."
// @Failure default {string} string "An error has ocurred"
// @Router /users/get_feed_requests_from_me [get]
func GetFeedRequestsFromMe(c *gin.Context) {
	allMyRequests, err := db.GetAllFeedRequested(IDUser)
	if err != nil {
		c.String(http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	var (
		info                     string
		UserWhoRequestedFeedInfo []string
	)

	for _, eachRequests := range allMyRequests {
		info = " " + eachRequests.RequestedUserName + " " + eachRequests.RequestedUserLastName + " requests your feedback."
		UserWhoRequestedFeedInfo = append(UserWhoRequestedFeedInfo, info)
	}

	c.JSON(http.StatusCreated, UserWhoRequestedFeedInfo)
}

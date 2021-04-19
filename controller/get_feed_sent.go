package controller

import (
	"fmt"
	"net/http"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/db"
	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"github.com/gin-gonic/gin"
)

// GetFeedsSent godoc
// @Description get feedback sent by id
// @id GetFeedsSent
// @Summary is used to get all the feedbacks of some user.
// @Param id path string true "Account ID"
// @Param Authorization header string true "jwt token"
// @Produce json
// @Success 201 {string} string "Feedbacks recovered successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred recovering the feedbacks."
// @Failure default {string} string "An error has ocurred"
// @Router /users/get_feedback_sent/{id} [get]
func GetFeedsSent(c *gin.Context) {
	id := c.Param("id")

	targetUser, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exist.")
		return
	}
	var UserAllFeedbacks []models.Feedback

	for _, feedback := range targetUser.FeedbackStatus.FeedbacksSended {
		feedResult, err := db.GetSelectedFeedBack(feedback)
		if err != nil {
			c.String(http.StatusBadRequest, "Feedback does not exist.")
			return
		}

		UserAllFeedbacks = append(UserAllFeedbacks, feedResult)
	}
	fmt.Println(UserAllFeedbacks)

	c.JSON(http.StatusCreated, UserAllFeedbacks)
}

package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

// Ban user godoc
// @Description it lets you ban a user if you're admin
// @id banuser
// @Summary is used to ban users.
// @Param Authorization header string true "jwt token"
// @Param id path string true "id"
// @Success 201 {string} string "User banned."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Unauthorized"
// @Failure 500 {string} string "An error has ocurred trying to ban the user."
// @Failure default {string} string "An error has ocurred"
// @Router /users/ban/{id} [patch]
func BanUser(c *gin.Context) {
	id := c.Param("id")

	if id == IDUser {
		c.String(http.StatusInternalServerError, "Are you trying to ban yourself? (...)")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusBadRequest, "User does not exist.")
		return
	}
	if user.Role == "admin" {
		c.String(http.StatusBadRequest, "You're trying to ban an admin.")
		return
	}
	if user.Enabled == false {
		c.String(http.StatusBadRequest, "User already banned.")
		return
	}
	var userEmpty models.User
	userEmpty.Enabled = false

	_, err = db.ModifyUser(userEmpty, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error trying to modify the user")
		return
	}

	c.String(http.StatusCreated, "User banned.")
}

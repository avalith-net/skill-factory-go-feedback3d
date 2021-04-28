package controller

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/gin-gonic/gin"
)

//GetProfilePhoto envía el avatar al http
func GetProfilePhoto(c *gin.Context) {

	id := c.Param("id")

	if len(id) < 1 {
		c.String(http.StatusBadRequest, "Bad request.")
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		c.String(http.StatusNotFound, "User Not Found.")
		return
	}

	if len(user.ProfilePicture) < 1 {
		c.String(http.StatusNotFound, "Image not found.")
		return
	}
	c.File("uploads/profilePicture/" + user.ProfilePicture)
}

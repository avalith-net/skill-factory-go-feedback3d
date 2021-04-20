package controller

import (
	"net/http"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/avalith-net/skill-factory-go-feedback3d/services"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {
	var user = models.User{}

	user.Name = c.Query("name")
	user.LastName = c.Query("lastname")
	user.Email = c.Query("email")
	file, _ := c.FormFile("photo")

	if structs.IsZero(user) {
		c.String(http.StatusBadRequest, "Must complete at least one field")
		return
	}
	//photo manager call
	var ext string
	if file != nil {
		ext, _ = services.ManagePhoto(file, IDUser)
		user.ProfilePicture = IDUser + "." + ext
	}

	ok, err := db.ModifyUser(user, IDUser)
	if !ok {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusCreated, "Success")
}

package controller

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

//SetProfilePicture is used to set the profile picture
func SetProfilePicture(c *gin.Context) {
	file, _ := c.FormFile("profilePicture")
	fileContent, _ := file.Open()
	var extension = strings.Split(file.Filename, ".")[1]

	// /* The profile picture is stored in "profilePicture" folder that is previously created to make sure
	// that everything is able to work : folder uploads and inside: folder profilePicture*/
	fProfilePicture := "uploads/profilePicture/" + IDUser + "." + extension

	f, err := os.OpenFile(fProfilePicture, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.String(http.StatusBadRequest, "Error setting account picture  "+err.Error())
		return
	}
	_, err = io.Copy(f, fileContent)
	if err != nil {
		c.String(http.StatusBadRequest, "Error trying to copy the picture  "+err.Error())
		return
	}

	/*recording the change in the database */
	var user models.User
	var status bool
	user.ProfilePicture = IDUser + "." + extension

	status, err = db.ModifyUser(user, IDUser)
	if err != nil || status == false {
		c.String(http.StatusBadRequest, "Database error  "+err.Error())
		return
	}

	c.String(http.StatusCreated, "Success")
}

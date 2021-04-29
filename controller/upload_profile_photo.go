package controller

import (
	"net/http"
	"strings"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/gin-gonic/gin"
)

// SetProfilePicture godoc
// @Description is used to change the account picture.
// @id setProfilePicture
// @Summary is used to change the account picture.
// @Accept  multipart/form-data
// @Produce  json
// @Param profilePicture formData file true "account image"
// @Param Authorization header string true "jwt token"
// @Success 201 {string} string "Profile picture set successfully."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Error setting account picture."
// @Failure 500 {string} string "Error trying to copy the picture."
// @Failure default {string} string "Database error"
// @Router /setProfilePic [post]
func SetProfilePicture(c *gin.Context) {
	file, err := c.FormFile("profilePicture")

	if file == nil || err != nil {
		c.String(http.StatusBadRequest, "Form err "+err.Error())
		return
	}

	extension := strings.Split(file.Filename, ".")[1]
	file.Filename = IDUser + "." + extension
	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "uploads/profilePicture/"+file.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not upload file."+err.Error())
		return
	}

	/*recording the change in the database */
	var user models.User
	var status bool
	user.ProfilePicture = file.Filename

	status, err = db.ModifyUser(user, IDUser)
	if err != nil || !status {
		c.String(http.StatusInternalServerError, "Database error  "+err.Error())
		return
	}

	c.String(http.StatusCreated, "Success")
}

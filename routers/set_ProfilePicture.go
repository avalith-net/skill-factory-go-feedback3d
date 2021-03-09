package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
)

//SetProfilePicture is used to set the profile picture
func SetProfilePicture(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("profilePicture")

	var extension = strings.Split(handler.Filename, ".")[1]

	/* The profile picture is stored in "profilePicture" folder that is previously created to make sure
	that everything is able to work : folder uploads and inside: folder profilePicture*/
	var fProfilePicture string = "uploads/profilePicture/" + IDUser + "." + extension

	f, err := os.OpenFile(fProfilePicture, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error setting account picture  "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error trying to copy the picture "+err.Error(), http.StatusBadRequest)
		return
	}

	/*recording the change in the database */
	var user models.User
	var status bool
	user.ProfilePicture = IDUser + "." + extension

	status, err = db.ModifyUser(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Database error "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusCreated)
}

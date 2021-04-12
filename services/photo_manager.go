package services

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func ManagePhoto(file *multipart.FileHeader, id string) (string, error) {
	fileContent, _ := file.Open()
	var extension = strings.Split(file.Filename, ".")[1]

	// /* The profile picture is stored in "profilePicture" folder that is previously created to make sure
	// that everything is able to work : folder uploads and inside: folder profilePicture*/
	fProfilePicture := "uploads/profilePicture/" + id + "." + extension

	f, err := os.OpenFile(fProfilePicture, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, fileContent)
	if err != nil {
		return "", err
	}

	return extension, nil
}

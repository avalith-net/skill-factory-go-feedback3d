package controller

import (
	"encoding/json"
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

// Registro godoc
// @Description Used to register to the app
// @id registro
// @Summary Used to register to the app
// @Param credentials body string true "Json body with name, lastname, email and password"
// @Accept  json
// @Success 201 {string} string "User created."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "internal error"
// @Failure 500 {string} string "An error has ocurred sending the email."
// @Failure default {string} string "An error has ocurred"
// @Router /sign_up [post]
func SignUp(c *gin.Context) {

	var t models.User
	err := json.NewDecoder(c.Request.Body).Decode(&t)
	if err != nil {
		c.String(http.StatusBadRequest, "Error checking the data: "+err.Error())
		return
	}

	//data validation
	if len(t.Email) == 0 {
		c.String(http.StatusBadRequest, "must add email.")
		return
	}
	if len(t.Password) < 6 {
		c.String(http.StatusBadRequest, "Your password must be at least 6 characters long.")
		return
	}
	_, encontrado, _ := db.UserAlreadyExist(t.Email)
	if encontrado == true {
		c.String(http.StatusBadRequest, "Email already registered.")
		return
	}
	_, status, err := db.AddRegister(t)
	if err != nil {
		c.String(http.StatusInternalServerError, "Database error "+err.Error())
		return
	}
	if status == false {
		c.String(http.StatusInternalServerError, "Error, Register not added.")
		return
	}
	c.String(http.StatusCreated, "Success")
}

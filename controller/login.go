package controller

import (
	"net/http"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/auth"
	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Description login to the app.
// @id login
// @Summary is used to login to the application.
// @Param credentials body string true "Json body with email and password"
// @Accept  json
// @Success 201 {string} string "Successful Login."
// @Header 201 {string} string "Status created"
// @Failure 400 {string} string "Wrong mail or password."
// @Failure 500 {string} string "Error generating the token."
// @Failure default {string} string "An error has ocurred"
// @Router /login [post]
func Login(c *gin.Context) {

	var usu models.User

	if err := c.ShouldBindJSON(&usu); err != nil {
		c.String(http.StatusBadRequest, "Check your form.")
		return
	}

	//email validation
	if len(usu.Email) == 0 {
		c.String(http.StatusBadRequest, "Email needed.")
		return
	}
	document, exists := db.LoginAttempt(usu.Email, usu.Password)
	if !exists {
		c.String(http.StatusBadRequest, "Wrong mail or password.")
		return
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	jwtKey, err := auth.GenerateJWT(document, expirationTime)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error generating the token "+err.Error())
		return
	}

	resp := models.LoginReply{
		Token: jwtKey,
	}

	c.JSON(http.StatusCreated, resp)

	//domain : Production and development. Add in .env file

	c.SetCookie("token", jwtKey, int(expirationTime.Unix()), "/", "localhost", false, true)
}

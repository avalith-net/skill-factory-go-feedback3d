package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/blotin1993/feedback-api/auth"
	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/models"
	"github.com/gin-gonic/gin"
)

//Login validation
func Login(c *gin.Context) {

	var usu models.User
	err := json.NewDecoder(c.Request.Body).Decode(&usu)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//email validation
	if len(usu.Email) == 0 {
		c.String(http.StatusBadRequest, "Email needed.")
		return
	}
	document, exists := db.LoginAttempt(usu.Email, usu.Password)
	if exists == false {
		c.String(http.StatusBadRequest, "Wrong user or password.")
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

	c.SetCookie("token", jwtKey, int(expirationTime.Unix()), "/", "localhost", false, true)
}

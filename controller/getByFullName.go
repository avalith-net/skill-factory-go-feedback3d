package controller

import (
	"net/http"
	"strings"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

//  godoc
// @Description
// @id
// @Summary
// @Router / []
func GetByFullName(c *gin.Context) {
	//parámetros: pag, agregar.
	fullName := c.Query("fullname")
	if len(fullName) < 1 {
		c.String(http.StatusBadRequest, "You must enter a name.")
		return
	}
	fullNameArray := strings.Fields(fullName)

	arrLen := len(fullNameArray)
	var name, lastName string

	switch arrLen {
	case 1:
		name, lastName = fullNameArray[0], fullNameArray[0]
	case 2:
		name = fullNameArray[0]
		lastName = fullNameArray[1]
	case 3:
		name = fullNameArray[0] + " " + fullNameArray[1]
		lastName = fullNameArray[2]
	case 4:
		name = fullNameArray[0] + " " + fullNameArray[1]
		lastName = fullNameArray[2] + " " + fullNameArray[3]
	}

	users, err := db.GetUserByFullName(name, lastName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading users..."+err.Error())
		return
	}

	c.JSON(http.StatusCreated, users)
}

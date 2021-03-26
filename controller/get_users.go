package controller

import (
	"net/http"

	"github.com/blotin1993/feedback-api/db"
	"github.com/gin-gonic/gin"
)

//  godoc
// @Description
// @id
// @Summary
// @Router / []
func GetUsers(c *gin.Context) {
	//parámetros: pag, agregar.
	users, err := db.GetUsersDb()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading users...")
	}

	c.JSON(http.StatusCreated, users)
}

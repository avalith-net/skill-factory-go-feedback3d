package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/router"
)

// @title feedback API
// @version 1.0
// @description Aplicación que permite realizar feedbacks entre los miembros de un equipo de trabajo.
// @termsOfService https://avalith.net/about-us/terms-of-use

// @contact.name Avalith
// @contact.url https://avalith.net/
// @contact.email vlotin_gaming@gmail.com

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {

	if db.CheckConnection() == 0 {

		log.Fatal("No connection to the BD")
		return
	}
	router.SetRoutes()
}

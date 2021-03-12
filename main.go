package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/handlers"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @host localhost:8080

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1
func main() {

	if db.CheckConnection() == 0 {

		log.Fatal("No connection to the BD")
		return
	}
	handlers.SetRoutes()
}

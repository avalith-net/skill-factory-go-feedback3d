package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/router"
)

func main() {

	if db.CheckConnection() == 0 {

		log.Fatal("No connection to the BD")
		return
	}
	router.SetRoutes()
}

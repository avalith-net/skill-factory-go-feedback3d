package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/blotin1993/feedback-api/db"
	"github.com/blotin1993/feedback-api/handlers"
)

func main() {

	/*pass, _ := db.PassEncrypt("nofede")
	fmt.Println(pass)
	dpass, _ := db.DecryptPassword(pass)
	fmt.Println(dpass)*/

	if db.CheckConnection() == 0 {

		log.Fatal("No connection to the BD")
		return
	}
	handlers.SetRoutes()
}

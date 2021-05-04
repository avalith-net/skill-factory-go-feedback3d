package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron/v3"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/router"
)

func main() {

	if db.CheckConnection() == 0 {

		log.Fatal("No connection to the BD")
		return
	}

	//watch for changes in db
	go db.WatchTimeLeft()
	//start cron scheduler
	c := cron.New()
	c.AddFunc("@midnight", func() {
		err := db.UpdateTimeLeft()
		if err != nil {
			log.Println(err)
			return
		}
	})
	c.Start()

	router.SetRoutes()
}

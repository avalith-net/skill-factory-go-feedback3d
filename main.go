package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/avalith-net/skill-factory-go-feedback3d/db"
	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/avalith-net/skill-factory-go-feedback3d/router"
)

func main() {

	if db.CheckConnection() == 0 {

		log.Fatal("No connection to the BD")
		return
	}

	t, err := db.GetTime()
	if err == nil {
		go test(time.Now(), t)
	} else {
		fmt.Println(err)
	}

	router.SetRoutes()
}

func test(start time.Time, reqTime []models.FeedbacksRequested) {
	for {
		//... operation that takes 2 seconds + 20 milliseconds ...
		for _, cur := range reqTime {
			try := cur.SentDate.Sub(start)
			switch {
			case try < -24*time.Hour:
				fmt.Println("Pasó más de un día")
				fmt.Println(cur.SentDate.Sub(start))
			case try < -360*time.Hour:
				fmt.Println("Pasó más de 15 días")
				// delete request.
			}
		}
		time.Sleep(120 * time.Hour)
	}
}

// func isTime(t *time.Time) bool {
// 	if t.Sub(time.Now()) < -24*time.Hour {
// 		*t = time.Now()
// 		return true
// 	}
// 	return true
// }

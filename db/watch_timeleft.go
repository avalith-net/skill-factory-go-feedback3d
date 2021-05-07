package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"github.com/avalith-net/skill-factory-go-feedback3d/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func WatchTimeLeft() {

	coll := MongoCN.Database("feedback-db").Collection("feedbacks-requested")

	pipeline := bson.D{
		{
			"$match", bson.D{
				{"operationType", "update"},
				{"updateDescription.updatedFields.timeleft", bson.D{
					{"$in", bson.A{0, 1, 5, 10}},
				}},
			},
		},
	}

	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	requestsStream, err := coll.Watch(context.TODO(), mongo.Pipeline{pipeline}, opts)
	if err != nil {
		panic(err)
	}

	for requestsStream.Next(context.TODO()) {
		var data models.ChangeEvent

		raw := requestsStream.Current
		err := bson.Unmarshal(raw, &data)
		if err != nil {
			return
		}
		if data.UpdateDescription.UpdatedFields.Timeleft == 0 {
			DeleteFeedbackRequested(data.FullDocument.FeedbackRequestID)
		} else {
			email, err := GetUserMailById(data.FullDocument.RequestedUserID)
			if err != nil {
				return
			}

			//TODO: Set scheduler for 6 am (if possible) done
			// send link attached in mail body* Done
			bodyString := "Hey!\nYou have " + fmt.Sprint(data.UpdateDescription.UpdatedFields.Timeleft) + " days left to make your feedback,\nHurry up!\n" +
				"Follow this link: http:localhost:8080/target_id=" + data.FullDocument.LoggedUserId

			time.AfterFunc(6*time.Hour, func() {
				services.SendEmail(email, "You have a pending feedback request.", bodyString)
				log.Println("Email sent by watch routine")
			})
		}
	}

}

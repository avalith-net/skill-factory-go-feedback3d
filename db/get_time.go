package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetUser .
func GetTime() ([]models.FeedbacksRequested, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")

	var fb []models.FeedbacksRequested

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var eachFb models.FeedbacksRequested
		err := cursor.Decode(&eachFb)
		if err != nil {
			return nil, err
		}
		fb = append(fb, eachFb)
	}
	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	cursor.Close(ctx)
	return fb, nil
}

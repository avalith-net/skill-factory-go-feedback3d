package db

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//DeleteFeedback delete the feed with the given feedback ID
func DeleteFeedback(feedbackID string) (*mongo.DeleteResult, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	var isDeleted bool

	objID, err := primitive.ObjectIDFromHex(feedbackID)
	if err != nil {
		err = errors.New("error finding the feed to delete")
	}

	condition := bson.M{
		"_id": objID,
	}

	deleteResult, err := col.DeleteOne(ctx, condition)
	if err != nil {
		isDeleted = false
		log.Fatal(err)
	}
	isDeleted = true

	return deleteResult, isDeleted
}

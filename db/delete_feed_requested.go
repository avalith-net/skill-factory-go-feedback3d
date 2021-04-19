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

func DeleteFeedbackRequested(feedbackID string) (*mongo.DeleteResult, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")

	var isDeleted bool

	objID, err := primitive.ObjectIDFromHex(feedbackID)
	if err != nil {
		err = errors.New("error finding the feedbackRequested to delete")
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

package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddFeedbackRequested(fbr models.FeedbacksRequested) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")
	result, err := col.InsertOne(ctx, fbr)
	if err != nil {
		return "Error trying to insert new feedbackRequest in the database.", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.Hex(), true, nil
}
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSelectedFeedBackRequestObj(FeedStatusID string) (models.FeedbacksRequested, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")

	var feedbackStatusObj models.FeedbacksRequested

	objID, _ := primitive.ObjectIDFromHex(FeedStatusID)
	condition := bson.M{
		"_id": objID,
	}

	if err := col.FindOne(ctx, condition).Decode(&feedbackStatusObj); err != nil {
		fmt.Println("feedbackrequest obj not found with given ID " + err.Error())
		return feedbackStatusObj, err
	}

	return feedbackStatusObj, nil
}

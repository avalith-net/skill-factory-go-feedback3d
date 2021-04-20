package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFeedStatus(ID string) (models.FeedbackStatus, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoCN.Database("feedback-db")
	col := database.Collection("feedbacks-status")

	var feedStatus models.FeedbackStatus

	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&feedStatus)
	if err != nil {
		return feedStatus, err
	}
	return feedStatus, nil
}

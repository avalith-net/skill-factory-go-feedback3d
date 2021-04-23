package db

import (
	"context"
	"fmt"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFeedBackRequestedID(targetUserID string, loggedUserID string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")

	var feedStatus models.FeedbacksRequested

	condition := bson.M{
		"requested_user_id": targetUserID,
		"user_logged_id":    loggedUserID,
	}

	err := col.FindOne(ctx, condition).Decode(&feedStatus)

	if err != nil {
		fmt.Println("feedbackRequest ojb not found with given ID inside GetFeedBackRequestedID function ->" + err.Error())
		return feedStatus.ID.Hex(), false
	}

	return feedStatus.ID.Hex(), true
}

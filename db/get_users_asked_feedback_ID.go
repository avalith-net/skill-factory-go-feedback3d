package db

import (
	"context"
	"fmt"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsersAskedFeedbackID(loggedUserID string, targetUserID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("users_asks_feedback")

	var feedStatus models.UsersAskedFeed

	condition := bson.M{
		"user_asking_feed_id":    targetUserID,
		"user_asked_for_feed_id": loggedUserID,
	}

	err := col.FindOne(ctx, condition).Decode(&feedStatus)
	if err != nil {
		fmt.Println("userAsksFeed obj not found with given ID inside GetUsersAskedFeedbackID function ->" + err.Error())
		return feedStatus.ID.Hex(), err
	}

	return feedStatus.ID.Hex(), nil
}

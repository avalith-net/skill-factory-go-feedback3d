package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyFeedStatus(fbs models.FeedbackStatus, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedback-status")

	feedStatus := make(map[string]interface{})

	feedStatus["user_id"] = fbs.UserID
	feedStatus["users_asked_feed"] = fbs.UsersAskedFeed
	feedStatus["feedsrequested"] = fbs.FeedbacksRequested
	feedStatus["feeds_sended"] = fbs.FeedbacksSended

	updtString := bson.M{
		"$set": feedStatus,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil

}

package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsersAskingForFeed(targetUserID string) ([]*models.UsersAskedFeed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users_asks_feedback")

	var feedsUsersAskedFor []*models.UsersAskedFeed

	condicion := bson.M{
		"user_asked_for_feed_id": bson.M{"$eq": targetUserID},
	}

	var page int64 = 1

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	cursor, err := col.Find(ctx, condicion, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var eachFeed models.UsersAskedFeed
		err := cursor.Decode(&eachFeed)
		if err != nil {
			return nil, err
		}
		feedsUsersAskedFor = append(feedsUsersAskedFor, &eachFeed)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	return feedsUsersAskedFor, nil
}

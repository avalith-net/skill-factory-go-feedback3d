package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllFeedRequested(targetUserID string) ([]*models.FeedbacksRequested, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")

	var (
		feedsRequested []*models.FeedbacksRequested
		page           int64
	)

	condition := bson.M{"user_logged_id": bson.M{"$eq": targetUserID}}

	page = 1

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	cursor, err := col.Find(ctx, condition, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var eachFeed models.FeedbacksRequested
		err := cursor.Decode(&eachFeed)
		if err != nil {
			return nil, err
		}
		feedsRequested = append(feedsRequested, &eachFeed)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	return feedsRequested, nil
}

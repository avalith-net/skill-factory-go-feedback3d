package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FeedIsNotDisplayable is used to set to false the feedback condition where it is displayable or not.
func FeedIsNotDisplayable(feedID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	reportedFb := make(map[string]interface{})

	reportedFb["is_displayable"] = false

	updateFeed := bson.M{
		"$set": reportedFb,
	}

	objID, _ := primitive.ObjectIDFromHex(feedID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateFeed)
	if err != nil {
		return false, err
	}
	return true, nil
}

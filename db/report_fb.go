package db

import (
	"context"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ReportFeedback is used to report a feedback as user.
func ReportFeedback(feed models.Feedback, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	reportedFb := make(map[string]interface{})

	reportedFb["is_reported"] = true
	reportedFb["is_approbed"] = false

	updateFeed := bson.M{
		"$set": reportedFb,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateFeed)
	if err != nil {
		return false, err
	}
	return true, nil
}

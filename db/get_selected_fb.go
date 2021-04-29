package db

import (
	"context"
	"fmt"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetSelectedFeedBack return a specific feedback with the given feed id
func GetSelectedFeedBack(FeedbackID string) (models.Feedback, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	var feedback models.Feedback

	objID, _ := primitive.ObjectIDFromHex(FeedbackID)
	condition := bson.M{
		"_id":            objID,
		"is_displayable": true,
	}

	if err := col.FindOne(ctx, condition).Decode(&feedback); err != nil {
		fmt.Println("feedback not found with given ID " + err.Error())
		return feedback, err
	}

	return feedback, nil
}

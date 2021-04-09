package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetFeedFromDb .
func GetSelectedFeedBack(FeedbackID string) (models.Feedback, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	var feedback models.Feedback

	// ObjectIDFromHex creates a new ObjectID from a hex string. It returns an error if the hex string is not a
	// valid ObjectID.
	objID, _ := primitive.ObjectIDFromHex(FeedbackID)
	condition := bson.M{
		"_id": objID,
	}
	if err := col.FindOne(ctx, condition).Decode(&feedback); err != nil {
		//err = errors.New("error finding the fb with given ID")
		fmt.Println("feedback not found with given ID " + err.Error())
		return feedback, err
	}

	fmt.Println("El feedback es ")
	fmt.Println(feedback)

	return feedback, nil
}

package db

import (
	"context"
	"errors"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetFeedFromDb .
func GetFeedFromDb(ID string) ([]models.Feedback, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	var feedSlice []models.Feedback

	condicion := bson.M{
		"receiver_id": ID,
	}

	cur, err := col.Find(ctx, condicion)
	if err != nil {
		err = errors.New("Error al buscar los elementos")
	}

	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem models.Feedback
		err := cur.Decode(&elem)
		if err != nil {
			err = errors.New("Error al buscar los elementos")
		}
		feedSlice = append(feedSlice, elem)
	}
	// Close the cursor once finished
	cur.Close(ctx)

	return feedSlice, err
}

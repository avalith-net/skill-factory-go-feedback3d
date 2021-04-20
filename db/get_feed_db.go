package db

import (
	"context"
	"errors"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetFeedFromDb . Contion @param {true for receiver and false for issuer}
func GetFeedFromDb(ID string, condition bool) ([]models.Feedback, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	var feedSlice []models.Feedback

	// TODO : Agregar filtro. Cant, etc.
	var filter bson.M

	//true for receiver and false for issuer
	if condition {
		filter = bson.M{
			"receiver_id": ID,
		}
	} else {
		filter = bson.M{
			"issuer_id": ID,
		}
	}

	cur, err := col.Find(ctx, filter)
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

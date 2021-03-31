package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFeedbacks(ID string) ([]*models.Feedback, []*models.Feedback, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks")

	findOptions := options.Find()
	//Setting 20 feedbacks per search
	findOptions.SetLimit(20)

	var issuer []*models.Feedback
	var receiver []*models.Feedback

	issuerID := bson.M{
		"issuer_id": ID,
	}

	receiverID := bson.M{
		"receiver_id": ID,
	}

	cur, err := col.Find(ctx, issuerID, findOptions)
	if err != nil {
		return nil, nil, err
	}

	cur2, err := col.Find(ctx, receiverID, findOptions)
	if err != nil {
		return nil, nil, err
	}

	// Finding multiple documents returns a cursor
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var fbIssuer models.Feedback
		var fbReceiver models.Feedback

		err := cur.Decode(&fbIssuer)
		if err != nil {
			return nil, nil, err
		}

		err = cur2.Decode(&fbReceiver)
		if err != nil {
			return nil, nil, err
		}

		issuer = append(issuer, &fbIssuer)
		receiver = append(receiver, &fbReceiver)
	}

	if err := cur.Err(); err != nil {
		return nil, nil, err
	}

	if err := cur2.Err(); err != nil {
		return nil, nil, err
	}

	// Closing the cursors
	cur.Close(ctx)
	cur2.Close(ctx)

	return issuer, receiver, nil

}

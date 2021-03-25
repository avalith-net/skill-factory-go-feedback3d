package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllUsers .
func GetUsersDb() ([]*models.ReturnUser, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	// Pass these options to the Find method
	findOptions := options.Find()
	//Corregir
	findOptions.SetLimit(20)
	// array in which you can store the decoded documents
	var results []*models.ReturnUser

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := col.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var user models.ReturnUser
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		results = append(results, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(ctx)

	return results, nil
}

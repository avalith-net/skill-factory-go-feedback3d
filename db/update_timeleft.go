package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//UpdateTimeLeft .
func UpdateTimeLeft() error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("feedbacks-requested")

	//TODO add gt > 0
	result, err := col.UpdateMany(
		ctx,
		bson.M{},
		bson.D{
			{"$inc", bson.D{{"timeleft", -1}}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

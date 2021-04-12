package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateGraphic -
func UpdateGraphic(g models.ReturnUser, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	register := make(map[string]interface{}) // this is the map used to update the db register.

	register["graphic"] = g.Graphic

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}

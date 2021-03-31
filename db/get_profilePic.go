package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetProfilePic(email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	var profilePic string

	condition := bson.M{"email": email}

	err := col.FindOne(ctx, condition).Decode(&profilePic)
	if err != nil {
		return profilePic, err
	}

	return profilePic, nil
}

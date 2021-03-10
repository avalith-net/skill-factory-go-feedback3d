package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserAlreadyExist checks if the user is inside the db.
func UserAlreadyExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	condition := bson.M{"email": email} // email filter
	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result) // if it exists it will be passed to result
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}

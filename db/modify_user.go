package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/auth"
	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModifyUser -> allows us to modify the register in the db.
func ModifyUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	register := make(map[string]interface{}) // this is the map used to update the db register.

	//Some validations...
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		register["lastname"] = u.LastName
	}
	if len(u.ProfilePicture) > 0 {
		register["profilePicture"] = u.ProfilePicture
	}
	if len(u.Password) > 0 {
		u.Password, _ = auth.PassEncrypt(u.Password)
		register["password"] = u.Password
	}

	register["feedsrequested"] = u.FeedbacksRequested
	register["users_asked_feed"] = u.UsersAskedFeed

	//Improve
	if len(u.Name) == 0 && len(u.LastName) == 0 && len(u.ProfilePicture) == 0 && len(u.Password) == 0 {
		register["enabled"] = u.Enabled
	}

	updtString := bson.M{
		"$set": register,
	}

	fmt.Println()

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}

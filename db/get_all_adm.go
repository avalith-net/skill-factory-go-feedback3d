package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// It return a slice of all the users with role = "admin"
func GetAllAdmins() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	var admins []*models.User

	condicion := bson.M{
		"role": bson.M{"$eq": "admin"},
	}

	var page int64 = 1

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	cursor, err := col.Find(ctx, condicion, findOptions)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var eachAdm models.User
		err := cursor.Decode(&eachAdm)
		if err != nil {
			return nil, err
		}
		admins = append(admins, &eachAdm)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	return admins, nil
}

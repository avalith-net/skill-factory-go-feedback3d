package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// It return a slice of all the users with role = "admin"
func GetAllAdmins(page int64) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	var admins []*models.User

	condicion := bson.M{
		"role": bson.M{"$eq": "admin"},
	}

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

	fmt.Println(admins)
	return admins, nil
}

package db

import (
	"context"
	"time"

	"github.com/avalith-net/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetUser .
func GetUserMailById(ID string) (string, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	var email models.JustEmail
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&email)
	if err != nil {
		return "", err
	}
	return email.Email, nil
}

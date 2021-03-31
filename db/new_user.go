package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/auth"
	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddRegister is used to store new users into the db.
func AddRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")
	var err error
	u.Password, err = auth.PassEncrypt(u.Password)
	u.Enabled = true
	u.Role = "user"
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

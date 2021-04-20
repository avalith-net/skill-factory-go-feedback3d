package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/auth"
	"github.com/JoaoPaulo87/skill-factory-go-feedback3d/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddRegister is used to store new users into the db.
func AddRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoCN.Database("feedback-db")
	col := database.Collection("users")

	var err error

	u.Password, err = auth.PassEncrypt(u.Password)
	u.Enabled = true
	u.Role = "user"

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		fmt.Println("Error trying to insert the register in database")
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.Hex(), true, nil
}

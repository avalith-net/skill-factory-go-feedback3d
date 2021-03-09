package db
 
import (
    "context"
    "fmt"
    "time"
 
    "github.com/blotin1993/feedback-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


func GetUser(ID string) (models.User, error) {
 
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()
    db := MongoCN.Database("feedback-api")
    col := db.Collection("users")
 
    var user models.User
    objID, _ := primitive.ObjectIDFromHex(ID)
 
    condicion := bson.M{
        "_id": objID,
    }
    err := col.FindOne(ctx, condicion).Decode(&user)
    user.Password = ""
    if err != nil {
        fmt.Println("User not found " + err.Error())
        return user, err
    }
    return user, nil
}

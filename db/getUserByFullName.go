package db

import (
	"context"
	"time"

	"github.com/blotin1993/feedback-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetUserByFullName .
func GetUserByFullName(name, lastName string) ([]*models.SearchUser, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("feedback-db")
	col := db.Collection("users")

	// Pass these options to the Find method
	findOptions := options.Find()
	//Corregir
	findOptions.SetLimit(20)

	filterOne := bson.D{
		{"$or",
			bson.A{
				bson.D{
					{"$or",
						bson.A{
							bson.D{
								{"$and",
									bson.A{
										bson.D{{"name", bson.M{"$regex": `(?i)` + name}}},
										bson.D{{"lastname", bson.M{"$regex": `(?i)` + lastName}}},
									},
								},
							},
							bson.D{
								{"$and",
									bson.A{
										bson.D{{"name", bson.M{"$regex": `(?i)` + lastName}}},
										bson.D{{"lastname", bson.M{"$regex": `(?i)` + name}}},
									},
								},
							},
						}},
				},
			},
		},
	}

	filterTwo := bson.D{
		{"$or",
			bson.A{
				bson.D{{"lastname", bson.M{"$regex": `(?i)` + lastName}}},
				bson.D{{"lastname", bson.M{"$regex": `(?i)` + name}}},
				bson.D{{"name", bson.M{"$regex": `(?i)` + name}}},
			}},
	}

	firstResult, err := customFind(col, ctx, filterOne, findOptions)
	if err != nil {
		return nil, err
	}
	if firstResult == nil {
		secondResult, err := customFind(col, ctx, filterTwo, findOptions)
		if err != nil {
			return nil, err
		}
		return secondResult, nil
	}
	return firstResult, nil
}

func customFind(col *mongo.Collection, ctx context.Context, filter primitive.D, findOptions *options.FindOptions) ([]*models.SearchUser, error) {
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := col.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var results []*models.SearchUser
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {

		// create a value into which the single document can be decoded
		var user models.SearchUser
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		results = append(results, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(ctx)
	return results, nil
}

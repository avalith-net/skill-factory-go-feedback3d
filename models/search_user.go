package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//SearchUser is the struct used to return when we search the user
type SearchUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	LastName string             `bson:"lastname" json:"lastname,omitempty"`
	Role     string             `bson:"role" json:"role"`
}

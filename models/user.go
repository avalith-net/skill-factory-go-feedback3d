package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User structure
type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name,omitempty"`
	LastName       string             `bson:"lastname" json:"lastname,omitempty"`
	Email          string             `bson:"email" json:"email" validate:"required,email"`
	Password       string             `bson:"password" json:"password,omitempty"`
	ProfilePicture string             `bson:"profilePicture" json:"profilePicture,omitempty"`
	Graphic        Graphic            `bson:"graphic" json:"graphic,omitempty"`
	Role           string             `bson:"role" json:"role"`
	Enabled        bool               `bson:"enabled" json:"enabled,omitempty"`
	FeedbackStatus FeedbackStatus     `bson:"feedbackstatus" json:"feedbackstatus,omitempty"`
}

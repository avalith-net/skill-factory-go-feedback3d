package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedbacksRequested struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RequestedUserID       string             `bson:"requested_user_id,omitempty" json:"requested_user_id,omitempty"`
	UserLoggedID          string             `bson:"user_logged_id,omitempty" json:"user_logged_id,omitempty"`
	RequestedUserName     string             `bson:"requested_user_name" json:"name,omitempty"`
	RequestedUserLastName string             `bson:"requested_user_lastname" json:"lastname,omitempty"`
	SentDate              time.Time          `bson:"sentdate" json:"sentdate,omitempty"`
	TimeLeft              int32              `bson:"timeleft" json:"timeleft,omitempty"`
}

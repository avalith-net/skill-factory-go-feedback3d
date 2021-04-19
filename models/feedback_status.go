package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeedbackStatus struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID             string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UsersAskedFeed     UsersAskedFeed     `bson:"users_asked_feed" json:"users_asked_feed,omitempty"`
	FeedbacksRequested FeedbacksRequested `bson:"feedsrequested" json:"feedsrequested,omitempty"`
	FeedbacksSended    []string           `bson:"feeds_sended" json:"feeds_sended,omitempty"`
}

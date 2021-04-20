package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersAskedFeed struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserWhoAskFeedID   string             `bson:"user_asking_feed_id,omitempty" json:"user_asks_id,omitempty"`
	UserAskedForFeedID string             `bson:"user_asked_for_feed_id,omitempty" json:"user_asking_feed_id,omitempty"`
	NameWhoAskFeed     string             `bson:"name_who_asked" json:"name,omitempty"`
	LastNameWhoAskFeed string             `bson:"lastname_who_asked" json:"lastname,omitempty"`
	SentDate           time.Time          `bson:"sentdate" json:"sentdate,omitempty"`
}

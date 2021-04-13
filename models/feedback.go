package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Feedback structure
type Feedback struct {
	IssuerID        string             `bson:"issuer_id,omitempty" json:"issuer_id" deepcopier:"skip"`
	ReceiverID      string             `bson:"receiver_id,omitempty" json:"receiver_id" deepcopier:"skip"`
	Date            time.Time          `bson:"date" json:"date,omitempty" deepcopier:"skip"`
	TechArea        TechArea           `bson:"techarea,omitempty,inline" json:"techarea,omitempty"`
	TeamArea        TeamArea           `bson:"teamarea,omitempty,inline" json:"teamarea,omitempty"`
	PerformanceArea PerformanceArea    `bson:"performancearea,omitempty,inline" json:"performancearea,omitempty"`
	Message         string             `bson:"message,omitempty" json:"message"`
	ID              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	IsDisabled      bool               `bson:"is_disabled" json:"is_disabled"`
	IsApprobed      bool               `bson:"is_approbed" json:"is_appobed"`
	IsReported      bool               `bson:"is_reported" json:"is_reported"`
}

//TechArea .
type TechArea struct {
	Message       string `bson:"tamessage,omitempty" json:"tamessage"`
	TechKnowledge string `bson:"techKnowledge,omitempty" json:"techknowledge"`
	BestPractices string `bson:"bestPractices,omitempty" json:"bestpractices"`
	CodingStyle   string `bson:"codingStyle,omitempty" json:"codingstyle"`
}

//TeamArea .
type TeamArea struct {
	Message       string `bson:"temessage,omitempty" json:"temessage"`
	TeamPlayer    string `bson:"teamPlayer,omitempty" json:"teamplayer"`
	Commited      string `bson:"commited,omitempty" json:"commited"`
	Communication string `bson:"communication,omitempty" json:"communication"`
}

//PerformanceArea .
type PerformanceArea struct {
	Message        string `bson:"pamessage,omitempty" json:"pamessage"`
	WorkQuality    string `bson:"workQuality,omitempty" json:"workquality"`
	ClientOriented string `bson:"clientOriented,omitempty" json:"clientoriented"`
}

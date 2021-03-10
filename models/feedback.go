package models

import (
	"time"
)

//Feedback .
type Feedback struct {
	IssuerID   string    `bson:"issuer_id,omitempty" json:"issuer_id"`
	ReceiverID string    `bson:"receiver_id,omitempty" json:"receiver_id"`
	Date       time.Time `bson:"date" json:"date,omitempty"`
	TechArea   struct {
		TechKnowledge string `bson:"techKnowledge,omitempty" json:"techknowledge"`
		BestPractices string `bson:"bestPractices,omitempty" json:"bestpractices"`
		CodingStyle   string `bson:"codingStyle,omitempty" json:"codingstyle"`
	}
	TeamArea struct {
		TeamPlayer    string `bson:"teamPlayer,omitempty" json:"teamplayer"`
		Commited      string `bson:"commited,omitempty" json:"commited"`
		Communication string `bson:"communication,omitempty" json:"communication"`
	}
	PerformanceArea struct {
		WorkQuality    string `bson:"workQuality,omitempty" json:"workquality"`
		ClientOriented string `bson:"clientOriented,omitempty" json:"clientoriented"`
	}

	Message string `bson:"message,omitempty" json:"message"`
}

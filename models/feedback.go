package models

import (
	"time"
)

//Feedback structure
type Feedback struct {
	IssuerID        string          `bson:"issuer_id,omitempty" json:"issuer_id" structs:"-"`
	ReceiverID      string          `bson:"receiver_id,omitempty" json:"receiver_id" structs:"-"`
	Date            time.Time       `bson:"date" json:"date,omitempty" structs:"-"`
	TechArea        TechArea        `bson:"techarea,omitempty,inline" json:"techarea,omitempty"`
	TeamArea        TeamArea        `bson:"teamarea,omitempty,inline" json:"teamarea,omitempty"`
	PerformanceArea PerformanceArea `bson:"performancearea,omitempty,inline" json:"performancearea,omitempty"`
	Message         string          `bson:"message,omitempty" json:"message,omitempty" validate:"omitempty,max=1500" structs:"-"`
	Skills          []Skill         `bson:"skills,omitempty" json:"skills,omitempty" structs:"-"`
}

//TechArea .
type TechArea struct {
	Message       string `bson:"tamessage,omitempty" json:"tamessage,omitempty" validate:"omitempty,max=500" structs:"-"`
	TechKnowledge string `bson:"techKnowledge,omitempty" json:"techknowledge,omitempty" validate:"omitempty,customoneof"`
	BestPractices string `bson:"bestPractices,omitempty" json:"bestpractices,omitempty" validate:"omitempty,customoneof"`
	CodingStyle   string `bson:"codingStyle,omitempty" json:"codingstyle,omitempty" validate:"omitempty,customoneof"`
}

//TeamArea .
type TeamArea struct {
	Message       string `bson:"temessage,omitempty" json:"temessage,omitempty" validate:"omitempty,max=500" structs:"-"`
	TeamPlayer    string `bson:"teamPlayer,omitempty" json:"teamplayer,omitempty" validate:"omitempty,customoneof"`
	Commited      string `bson:"commited,omitempty" json:"commited,omitempty" validate:"omitempty,customoneof"`
	Communication string `bson:"communication,omitempty" json:"communication,omitempty" validate:"omitempty,customoneof"`
}

//PerformanceArea .
type PerformanceArea struct {
	Message        string `bson:"pamessage,omitempty" json:"pamessage,omitempty" validate:"omitempty,max=500" structs:"-"`
	WorkQuality    string `bson:"workQuality,omitempty" json:"workquality,omitempty" validate:"omitempty,customoneof"`
	ClientOriented string `bson:"clientOriented,omitempty" json:"clientoriented,omitempty" validate:"omitempty,customoneof"`
}

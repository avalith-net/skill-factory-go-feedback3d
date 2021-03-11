package models

//FeedbackRaw structure
type FeedbackRaw struct {
	TechArea        TechArea        `bson:"techarea,omitempty,inline" json:"techarea,omitempty"`
	TeamArea        TeamArea        `bson:"teamarea,omitempty,inline" json:"teamarea,omitempty"`
	PerformanceArea PerformanceArea `bson:"performancearea,omitempty,inline" json:"performancearea,omitempty" `
	Message         string          `bson:"message,omitempty" json:"message"`
}

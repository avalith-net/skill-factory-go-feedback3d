package models

//FeedbackRaw structure
type FeedbackRaw struct {
	TechArea struct {
		Message       string `bson:"tamessage,omitempty" json:"tamessage"`
		TechKnowledge string `bson:"techKnowledge,omitempty" json:"techknowledge"`
		BestPractices string `bson:"bestPractices,omitempty" json:"bestpractices"`
		CodingStyle   string `bson:"codingStyle,omitempty" json:"codingstyle"`
	}
	TeamArea struct {
		Message       string `bson:"temessage,omitempty" json:"temessage"`
		TeamPlayer    string `bson:"teamPlayer,omitempty" json:"teamplayer"`
		Commited      string `bson:"commited,omitempty" json:"commited"`
		Communication string `bson:"communication,omitempty" json:"communication"`
	}
	PerformanceArea struct {
		Message        string `bson:"pamessage,omitempty" json:"pamessage"`
		WorkQuality    string `bson:"workQuality,omitempty" json:"workquality"`
		ClientOriented string `bson:"clientOriented,omitempty" json:"clientoriented"`
	}
	Message struct {
		Description string `bson:"message,omitempty" json:"message"`
	}
}

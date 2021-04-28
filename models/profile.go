package models

//User structure
type GeneralProfile struct {
	UserID                string  `bson:"user_id" json:"user_id,omitempty"`
	CompleteName          string  `bson:"completeName" json:"completeName,omitempty"`
	ProfilePicture        string  `bson:"profilePicture" json:"profilePicture,omitempty"`
	Role                  string  `bson:"role" json:"role,omitempty"`
	Graphic               Graphic `bson:"graphic" json:"graphic,omitempty"`
	FbIssuer              int     `bson:"FbIssuer" json:"FbIssuer,omitempty"`
	FbReceiver            int     `bson:"FbReceived" json:"FbReceived,omitempty"`
	FeedbackSent          int     `bson:"feedbackSent" json:"feedbackSent,omitempty"`
	FeedbacksRequested    int     `bson:"feedrequested" json:"feedrequested,omitempty"`
	FeedbackAskedForUsers int     `bson:"feedaskedforusers" json:"feedaskedforusers,omitempty"`
}

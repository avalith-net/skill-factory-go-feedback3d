package models

//User structure
type GeneralProfile struct {
	CompleteName       string   `bson:"completeName" json:"completeName,omitempty"`
	ProfilePicture     string   `bson:"profilePicture" json:"profilePicture,omitempty"`
	Role               string   `bson:"role" json:"role,omitempty"`
	FbIssuer           int      `bson:"FbIssuer" json:"FbIssuer,omitempty"`
	FbReceiver         int      `bson:"FbReceived" json:"FbReceived,omitempty"`
	UsersAskedFeed     []string `bson:"users_asked_feed" json:"users_asked_feed,omitempty"`
	FeedbacksRequested []string `bson:"feedsrequested" json:"feedsrequested,omitempty"`
}

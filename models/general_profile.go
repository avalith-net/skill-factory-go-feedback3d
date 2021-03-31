package models

//User structure
type GeneralProfile struct {
	CompleteName   string   `bson:"completeName" json:"completeName,omitempty"`
	ProfilePicture string   `bson:"profilePicture" json:"profilePicture,omitempty"`
	Role           string   `bson:"role" json:"role,omitempty"`
	Metrics        []string `bson:"metrics" json:"metrics"`
	FbIssuer       int      `bson:"FbIssuer" json:"FbIssuer,omitempty"`
	FbReceiver     int      `bson:"FbReceived" json:"FbReceived,omitempty"`
}

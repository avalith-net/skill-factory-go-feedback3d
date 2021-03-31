package models

//ReturnUser is the struct used to parse the user...
type ReturnUser struct {
	Name           string `bson:"name" json:"name,omitempty"`
	LastName       string `bson:"lastname" json:"lastname,omitempty"`
	Email          string `bson:"email" json:"email"`
	ProfilePicture string `bson:"profilePicture" json:"profilePicture,omitempty"`
	Enabled        bool   `bson:"enabled" json:"enabled,omitempty"`
	Role           string `bson:"role" json:"role"`
}

package models

//ReturnUser is the struct used to parse the user...
type ReturnUser struct {
	Enabled  bool   `bson:"enabled" json:"enabled"`
	Name     string `bson:"name" json:"name,omitempty"`
	LastName string `bson:"lastname" json:"lastname,omitempty"`
	Email    string `bson:"email" json:"email"`
	Role     string `bson:"role" json:"role"`
}

package models

type ReturnUser struct {
    Name           string             `bson:"name" json:"name,omitempty"`
	LastName       string             `bson:"lastname" json:"lastname,omitempty"`
	Email          string             `bson:"email" json:"email"`
}



package models

//SearchUser is the struct used to return when we search the user
type JustEmail struct {
	Email string `bson:"email" json:"email,omitempty"`
}

package models

//User structure
type AdminProfile struct {
	Profile GeneralProfile
	Metrics []Feedback `bson:"metrics" json:"metrics"`
}

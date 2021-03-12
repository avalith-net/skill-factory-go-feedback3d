package models

/*FeedbackRequest struct */
type FeedbackRequest struct {
	ReceiverID    string `bson:"receiver_id,omitempty" json:"receiver_id" deepcopier:"skip"`
	ReceiverEmail string `bson:"receiver_email,omitempty" json:"receiver_email" deepcopier:"skip"`
}

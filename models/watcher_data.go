package models

// ChangeEvent is the customized representation of a MongoDB change stream event that is captured and processed by
// this application.
type ChangeEvent struct {
	FullDocument struct {
		FeedbackRequestID string `bson:"_id"`
		LoggedUserId      string `bson:"user_logged_id"`
		RequestedUserID   string `bson:"requested_user_id"`
	} `bson:"fullDocument"`
	UpdateDescription struct {
		UpdatedFields struct {
			Timeleft int `bson:"timeleft"`
		} `bson:"updatedFields"`
	} `bson:"updateDescription"`
}

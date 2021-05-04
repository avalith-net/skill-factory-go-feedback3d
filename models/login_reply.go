package models

//LoginReply model used to answer login attempt.
type LoginReply struct {
	Token      string `json:"token,omitempty"`
	UserLogged User   `json:"user_logged,omitempty"`
}

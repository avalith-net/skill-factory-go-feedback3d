package models

//LoginReply ..
type LoginReply struct {
	// cuando colocamos el omitempty es porque en caso de error debe devolver vacío
	Token string `json:"token,omitempty"`
}

package data

type Tweet struct {
	Id      int32  `json:"id"`
	Text    string `json:"text" validate:"required"`
	Picture string `json:"picture"`
}

package data

type Tweet struct {
	Id      string `json:"id"`
	Text    string `json:"text" validate:"required"`
	Picture string `json:"picture"`
	UserId  string `json:"user_id"`
}
type TokenStr struct {
	Token string `json:"token"`
}

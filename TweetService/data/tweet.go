package data

type Tweet struct {
	Id        string `json:"id"`
	Text      string `json:"text" validate:"required"`
	Picture   string `json:"picture"`
	UserEmail string `json:"user_email"`
}
type TokenStr struct {
	Token string `json:"token"`
}

package data

type Tweet struct {
	Id       string `json:"id"`
	Text     string `json:"text" validate:"required"`
	Username string `json:"username"`
}
type TokenStr struct {
	Token string `json:"token"`
}

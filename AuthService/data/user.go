package data

type User struct {
	Username    string `json:"username" `
	Email       string `json:"email" `
	Password    string `json:"password"`
	Role        string `json:"-"`
	IsActivated bool
}

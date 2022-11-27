package data

type User struct {
	Username  string `json:"username" `
	Email     string `json:"email" `
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	Role      string `json:"-"`
}

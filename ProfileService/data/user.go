package data

type User struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Country        string `json:"country"`
	Age            int    `json:"age"`
	CompanyName    string `json:"company_name"`
	CompanyWebsite string `json:"company_website"`
	Private        bool   `json:"private"`
}

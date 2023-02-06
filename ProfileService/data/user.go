package data

type User struct {
	Username       string `json:"username"`
	FirstName      string `bson:"first-name,omitempty" json:"first_name"`
	LastName       string `bson:"last-name,omitempty" json:"last_name"`
	Email          string `json:"email"`
	Gender         string `bson:"gender,omitempty" json:"gender"`
	Country        string `bson:"country,omitempty" json:"country"`
	Age            int    `bson:"age,omitempty" json:"age"`
	CompanyName    string `bson:"company_name,omitempty" json:"company_name"`
	CompanyWebsite string `bson:"company_website,omitempty" json:"company_website"`
	Private        bool   `json:"private"`
	Role           string `json:"role"`
}

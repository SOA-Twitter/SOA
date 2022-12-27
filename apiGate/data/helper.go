package data

import (
	"encoding/json"
	"io"
)

type User struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Gender         Gender `json:"gender"`
	Country        string `json:"country"`
	Age            int    `json:"age"`
	CompanyName    string `json:"company_name"`
	CompanyWebsite string `json:"company_website"`
	Role           Role   `json:"role"`
	Private        bool   `json:"private"`
}
type UserInfo struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Gender         Gender `json:"gender"`
	Country        string `json:"country"`
	Age            int    `json:"age"`
	CompanyName    string `json:"company_name"`
	CompanyWebsite string `json:"company_website"`
	Private        bool   `json:"private"`
	Role           string `json:"role"`
}
type UserIsFollowed struct {
	IsFollowed bool `json:"is_followed"`
}
type UserNode struct {
	Username string `json:"username"`
}

type ChangePass struct {
	OldPassword      string `json:"old_password"`
	NewPassword      string `json:"new_password"`
	RepeatedPassword string `json:"repeated_password"`
}

type ManagePrivacy struct {
	Private bool `json:"private"`
}

type LikedTweet struct {
	Liked bool `json:"liked"`
}

type Email struct {
	Email string `json:"email"`
}

type RecoverProfile struct {
	NewPassword      string `json:"new_password"`
	RepeatedPassword string `json:"repeated_password"`
	RecoveryUUID     string `json:"recovery_uuid"`
}

type Tweet struct {
	Id       string `json:"id"`
	Text     string `json:"text" validate:"required"`
	Username string `json:"username"`
}

type Tweets struct {
	Tweets []*Tweet
}

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

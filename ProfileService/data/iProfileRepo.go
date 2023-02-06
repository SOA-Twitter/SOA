package data

type ProfileRepoInt interface {
	Register(user *User) error
	GetUserProfile(Username string) (*User, error)
}

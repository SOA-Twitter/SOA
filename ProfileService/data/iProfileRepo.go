package data

type ProfileRepoInt interface {
	Register(user *User) error
	GetUserProfile(userID string) (User, error)
}

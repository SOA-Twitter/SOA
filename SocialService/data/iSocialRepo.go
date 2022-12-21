package data

type SocialRepo interface {
	RegUser(username string) error
	Follow(username string, isPrivate bool) (string, error)
}

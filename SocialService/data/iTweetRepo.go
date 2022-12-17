package data

type SocialRepo interface {
	AddUser(username string) error
}

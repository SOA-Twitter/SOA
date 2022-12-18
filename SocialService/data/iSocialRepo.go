package data

type SocialRepo interface {
	RegUser(username string) error
}

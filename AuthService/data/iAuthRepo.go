package data

type AuthRepo interface {
	Register(us *User) error
	CheckCredentials(us string, pas string) error
	FindUserID(us string) (string, error)
}

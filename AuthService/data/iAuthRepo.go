package data

type AuthRepo interface {
	Register(us *User) error
	CheckCredentials(em string, pas string) error
	FindUserID(us string) (string, error)
}

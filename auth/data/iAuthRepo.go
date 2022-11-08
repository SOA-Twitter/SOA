package data

type AuthRepo interface {
	Register(us *User) error
	FindUser(us string, pas string) error
}

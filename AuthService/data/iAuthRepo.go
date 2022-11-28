package data

type AuthRepo interface {
	Register(us *User) error
	Delete(email string) error
	CheckCredentials(em string, pas string) error
	FindUserEmail(email string) (string, string, error)
}

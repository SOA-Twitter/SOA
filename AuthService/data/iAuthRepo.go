package data

type AuthRepo interface {
	Register(us *User) error
	Edit(email string) error
	Delete(email string) error
	CheckCredentials(em string, pas string) error
	FindUserEmail(email string) (string, string, error)
	FindUser(email string) (*User, error)
	SaveActivationRequest(activationUUID string, registeredEmail string) error
	FindActivationRequest(activationUUID string) (*ActivationRequest, error)
	DeleteActivationRequest(activationUUID string, email string) error
}

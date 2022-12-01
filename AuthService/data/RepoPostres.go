package data

import (
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthRepoPostgres struct {
	l  *log.Logger
	db *gorm.DB
}

func PostgresConnection(l *log.Logger) (*AuthRepoPostgres, error) {
	USERNAME := os.Getenv("USER")
	DB_HOST := os.Getenv("HOST")
	PASSWORD := os.Getenv("PASSWORD")
	DB_NAME := os.Getenv("DB")
	PORT := os.Getenv("PORT")
	l.Println("\n" + USERNAME + DB_HOST + PASSWORD + DB_NAME + PORT + "\n")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", DB_HOST, USERNAME, DB_NAME, PASSWORD, PORT)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		l.Println("Error establishing a database connection")
		return &AuthRepoPostgres{}, err
	}
	setup(db)
	l.Println("Successfully connected to postgres database")
	return &AuthRepoPostgres{l, db}, nil
}
func setup(db *gorm.DB) {
	db.AutoMigrate(&User{})
	//	TODO* check db tables
	db.AutoMigrate(&ActivationRequest{})
	db.AutoMigrate(&RecoveryRequest{})
}

func QueryError(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (ps *AuthRepoPostgres) Register(user *User) error {
	ps.l.Println("{AuthRepoPostgres} - Creating user")
	createdUser := ps.db.Create(user)
	var errMessage = createdUser.Error
	if createdUser.Error != nil {
		fmt.Println(errMessage)
		ps.l.Println("Unable to Create user.", errMessage)
		return QueryError("Please try again later.")
	}
	return nil
}

func (ps *AuthRepoPostgres) Edit(email string) error {
	ps.l.Println("{AuthRepoPostgres} - Editing user")
	err := ps.db.Model(&User{}).Where("Email = ?", email).Update("is_activated", true).Error
	if err != nil {
		return QueryError("Error updating user " + email)
	}
	return err
}
func (ps *AuthRepoPostgres) ChangePassword(email, password string) error {
	ps.l.Println("{AuthRepoPostgres} - Change password")
	err := ps.db.Model(&User{}).Where("Email = ?", email).Update("password", password).Error
	if err != nil {
		return QueryError("Error changing password")
	}
	return err
}

func (ps *AuthRepoPostgres) Delete(email string) error {
	ps.l.Println("{AuthRepoPostgres} - Delete user")
	user := &User{}
	err := ps.db.Where("Email = ?", email).Delete(&user).Error
	if err != nil {
		ps.l.Println("Error deleting user with email ")
		ps.l.Println(email)
		return QueryError("Error deleting user")
	}
	return nil
}

func (ps *AuthRepoPostgres) CheckCredentials(email string, password string) error {
	ps.l.Println("{AuthRepoPostgres} - Check if credentials are valid")
	user := &User{}
	if err := ps.db.Where("Email = ?", email).First(user).Error; err != nil {
		ps.l.Println("Invalid Email")
		return QueryError("Invalid credentials!!")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ps.l.Println("Invalid Password")
		return QueryError("Invalid credentials!!")
	}
	// TODO
	if user.IsActivated != true {
		return errors.New("account activation is needed before 1st login")
	}

	return nil
}

func (ps *AuthRepoPostgres) FindUserEmail(email string) (string, string, error) {
	ps.l.Println("{AuthRepoPostgres} - Find User Email")
	user := &User{}
	err := ps.db.Where("Email = ?", email).First(user).Error
	return user.Email, user.Role, err
}

func (ps *AuthRepoPostgres) FindUser(email string) (*User, error) {
	ps.l.Println("{AuthRepoPostgres} - Find Whole User")
	user := &User{}
	err := ps.db.Where("Email = ?", email).First(user).Error
	return user, err
}

func (ps *AuthRepoPostgres) SaveActivationRequest(activationUUID string, registeredEmail string) error {
	ps.l.Println("{AuthRepoPostgres} - Save acc. activation request")

	activationRequest := &ActivationRequest{
		ActivationUUID: activationUUID,
		Email:          registeredEmail,
	}
	//*TODO: check if db.Create() makes another table for new Struct, or whether it tries saving in Users table
	createdRequest := ps.db.Create(activationRequest)
	var errMessage = createdRequest.Error
	if createdRequest.Error != nil {
		fmt.Println(errMessage)
		ps.l.Println("Unable to Create Account Activation Request.", errMessage)
		return QueryError("Please try again later.")
	}
	return nil
}

func (ps *AuthRepoPostgres) FindActivationRequest(activationUUID string) (*ActivationRequest, error) {
	ps.l.Println("{AuthRepoPostgres} - Find Account Activation Request")
	activationReq := &ActivationRequest{}
	ps.l.Println("--------------------------")
	ps.l.Println(activationUUID)
	ps.l.Println("--------------------------")
	err := ps.db.Where("activation_uuid = ?", activationUUID).First(activationReq).Error
	return activationReq, err
}

func (ps *AuthRepoPostgres) DeleteActivationRequest(activationUUID string, email string) error {
	ps.l.Println("{AuthRepoPostgres} - Delete Account Activation Request")
	activationReq := &ActivationRequest{}
	err := ps.db.Where("activation_uuid = ? AND Email = ?", activationUUID, email).Delete(&activationReq).Error
	if err != nil {
		ps.l.Println("Error deleting Account Activation Request")
		ps.l.Println(email)
		return QueryError("Error deleting Account Activation Request")
	}
	return nil
}

func (ps *AuthRepoPostgres) SaveRecoveryRequest(recoveryUUID string, registeredEmail string) error {
	ps.l.Println("{AuthRepoPostgres} - Save Password Recovery request")

	recoveryRequest := &RecoveryRequest{
		RecoveryUUID: recoveryUUID,
		Email:        registeredEmail,
	}
	//*TODO: check if db.Create() makes another table for new Struct, or whether it tries saving in Users table
	createdRequest := ps.db.Create(recoveryRequest)
	var errMessage = createdRequest.Error
	if createdRequest.Error != nil {
		fmt.Println(errMessage)
		ps.l.Println("Unable to Create Password Recovery Request.", errMessage)
		return QueryError("Please try again later.")
	}
	return nil
}

func (ps *AuthRepoPostgres) FindRecoveryRequest(recoveryUUID string) (*RecoveryRequest, error) {
	ps.l.Println("{AuthRepoPostgres} - Find Password Recovery Request")
	recoveryRequest := &RecoveryRequest{}
	ps.l.Println("--------------------------")
	ps.l.Println(recoveryUUID)
	ps.l.Println("--------------------------")
	err := ps.db.Where("recovery_uuid = ?", recoveryUUID).First(recoveryRequest).Error
	return recoveryRequest, err
}

func (ps *AuthRepoPostgres) DeleteRecoveryRequest(recoveryUUID string, email string) error {
	ps.l.Println("{AuthRepoPostgres} - Delete Password Recovery Request")
	recoveryReq := &RecoveryRequest{}
	err := ps.db.Where("recovery_uuid = ? AND Email = ?", recoveryUUID, email).Delete(&recoveryReq).Error
	if err != nil {
		ps.l.Println("Error deleting Password Recovery Request")
		ps.l.Println(email)
		return QueryError("Error deleting Password Recovery Request")
	}
	return nil
}

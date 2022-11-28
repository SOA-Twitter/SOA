package data

import (
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

	return nil
}

func (ps *AuthRepoPostgres) FindUserEmail(email string) (string, string, error) {
	ps.l.Println("{AuthRepoPostgres} - Find User Email")
	user := &User{}
	err := ps.db.Where("Email = ?", email).First(user).Error
	return user.Email, user.Role, err
}

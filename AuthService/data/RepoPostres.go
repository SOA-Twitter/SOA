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
	// err := godotenv.Load("local.env")
	// if err != nil {
	// 	l.Fatalf("Some error occurred. Err: %s", err)
	// }
	USERNAME := os.Getenv("POSTGRES_USER")
	DB_HOST := os.Getenv("POSTGRES_HOST")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_NAME := os.Getenv("POSTGRES_DB")
	PORT := os.Getenv("POSTGRES_PORT")
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
func (ps *AuthRepoPostgres) CheckCredentials(username string, password string) error {
	ps.l.Println("{AuthRepoPostgres} - Check if credentials are valid")
	user := &User{}
	if err := ps.db.Where("Username = ?", username).First(user).Error; err != nil {
		ps.l.Println("Invalid Username")
		return QueryError("Invalid credentials!!")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ps.l.Println("Invalid Password")
		return QueryError("Invalid credentials!!")
	}

	return nil
}
func (ps *AuthRepoPostgres) FindUserID(username string) (string, error) {
	ps.l.Println("{AuthRepoPostgres} - Find User Id")
	user := &User{}
	err := ps.db.Where("Username = ?", username).First(user).Error
	return user.UserId, err

}

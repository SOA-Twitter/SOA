package data

import (
	"TweeterMicro/auth/model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

//func getEnv(key, fallback string) string {
//	value := os.Getenv(key)
//	if len(value) == 0 {
//		return fallback
//	}
//	return value
//}

func ConnectToDB() *gorm.DB {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	USERNAME := os.Getenv("db_username")
	DB_HOST := os.Getenv("db_host")
	PASSWORD := os.Getenv("db_password")
	DB_NAME := os.Getenv("db_name")
	PORT := os.Getenv("db_port")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", DB_HOST, USERNAME, DB_NAME, PASSWORD, PORT)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	fmt.Println("Successfully connected", db)
	return db
}

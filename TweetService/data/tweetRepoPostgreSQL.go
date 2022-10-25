package data

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//New implementation that uses postgres, it implements the iTweetRepo
type TweetRepoPostgreSql struct {
	log *log.Logger
	db  *gorm.DB
}

// Constructor
func NewPostgreSql(log *log.Logger) (TweetRepoPostgreSql, error) {
	username := os.Getenv("db_username")
	host := os.Getenv("db_host")
	password := os.Getenv("db_password")
	name := os.Getenv("db_name")
	port := os.Getenv("db_port")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", host, username, name, password, port)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		return TweetRepoPostgreSql{}, err
	}

	setup(db)
	return TweetRepoPostgreSql{log, db}, nil
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&Tweet{})
}

func (twRepo *TweetRepoPostgreSql) GetAll() Tweets {
	twRepo.log.Println("{TweetRepoPostgreSql} - getting all tweets")
	var tweets []*Tweet

	twRepo.db.Find(&tweets)

	return tweets
}

// func (twRepo *TweetRepoPostgreSql) GetTweets() Tweets {
// 	twRepo.log.Println("{TweetRepoPostgreSql} - getting tweets")
// 	var tweets []*Tweet

// 	twRepo.db.Where("deleted_on = ?", "").Find(&tweets)

// 	return tweets
// }

func (twRepo *TweetRepoPostgreSql) CreateTweet(tw *Tweet) {
	twRepo.log.Println("{TweetRepoPostgreSql} - posting tweet")

	// *TODO: ID auto-increment for new Tweets
	// tw.ID = getNextID()
	tw.CreatedOn = time.Now().UTC().String()
	tw.DeletedOn = ""

	twRepo.db.Create(tw)
}

func (twRepo *TweetRepoPostgreSql) PutTweet(tw *Tweet, id int) error {
	twRepo.log.Println("{TweetRepoPostgreSql} - putting tweet")

	var tweet Tweet

	twRepo.db.Where("id = ? AND deleted_on = ?", id, "").Find(&tweet)

	if tweet.ID == 0 {
		return errors.New(fmt.Sprintf("Tweet with id %d not found", id))
	}

	tweet.Text = tw.Text
	tweet.Picture = tw.Picture

	twRepo.db.Save(&tweet)

	*tw = tweet
	return nil
}

func (twRepo *TweetRepoPostgreSql) DeleteTweet(id int) error {
	twRepo.log.Println("{TweetRepoPostgreSql} - deleting tweet")

	var tweet Tweet

	twRepo.db.Where("id = ? AND deleted_on = ?", id, "").Find(&tweet)

	if tweet.ID == 0 {
		return errors.New(fmt.Sprintf("Tweet with id %d not found", id))
	}

	tweet.DeletedOn = time.Now().UTC().String()

	twRepo.db.Save(&tweet)
	return nil
}

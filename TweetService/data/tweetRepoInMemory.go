package data

import (
	"log"
	"time"
)

type Tweets []*Tweet

//New implementation that uses in-memory list, it implements the iTweetRepo
type TweetRepoInMemory struct {
	log *log.Logger
}

// Constructor
func NewInMemory(log *log.Logger) (TweetRepoInMemory, error) {
	return TweetRepoInMemory{log}, nil
}

func (twRepo *TweetRepoInMemory) GetAll() Tweets {
	twRepo.log.Println("{TweetRepoPostgreSql} - getting all tweets")
	return tweetList
}
func (twRepo *TweetRepoInMemory) CreateTweet(t *Tweet) {
	twRepo.log.Println("{TweetRepoPostgreSql} - posting tweet")
	t.ID = twRepo.getNextID()
	t.CreatedOn = time.Now().UTC().String()
	tweetList = append(tweetList, t)
}
func (twRepo *TweetRepoInMemory) PutTweet(tw *Tweet, id int) error {
	twRepo.log.Println("{TweetRepoPostgreSql} - putting tweet/ NOT IMPLEMENTED")
	return nil
}
func (twRepo *TweetRepoInMemory) DeleteTweet(id int) error {
	twRepo.log.Println("{TweetRepoPostgreSql} - deleting tweet/ NOT IMPLEMENTED")
	return nil
}

func (twRepo *TweetRepoInMemory) getNextID() int {
	max := 0
	for _, currTweet := range twRepo.GetAll() {
		if currTweet.ID > max {
			max = currTweet.ID
		}
	}
	return max + 1
}

var tweetList = []*Tweet{
	&Tweet{
		ID:        1,
		Text:      "Super cute cats",
		CreatedOn: time.Now().UTC().String(),
	},
	&Tweet{
		ID:        2,
		Text:      "Super cute dogs",
		CreatedOn: time.Now().UTC().String(),
	},
}

package data

import "time"

type Tweets []*Tweet

func GetAll() Tweets {
	return tweetList
}
func CreateTweet(t *Tweet) {
	t.ID = getNextID()
	t.CreatedOn = time.Now().UTC().String()
	tweetList = append(tweetList, t)
}

func getNextID() int {
	max := 0
	for _, currTweet := range GetAll() {
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

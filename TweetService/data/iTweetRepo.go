package data

import "TweetService/proto/tweet"

type TweetRepo interface {
	CreateTweet(tw *Tweet) error
	PutTweet(tw *Tweet, id int) error
	GetTweetsByUsername(username string) ([]*tweet.Tweet, error)
	DeleteTweet(id int) error
}

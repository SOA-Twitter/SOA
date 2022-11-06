package data

import "TweeterMicro/TweetService/proto/tweet"

type TweetRepo interface {
	GetAll() []*tweet.Tweet
	CreateTweet(tw *Tweet) error
	PutTweet(tw *Tweet, id int) error
	DeleteTweet(id int) error
}

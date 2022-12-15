package data

import "TweetService/proto/tweet"

type TweetRepo interface {
	CreateTweet(tw *Tweet) error
	GetTweetsByUsername(username string) ([]*tweet.Tweet, error)
	LikeTweet(id string, username string, like bool) error
	GetLikesByTweetId(id string) ([]*Like, error)
	GetLikesByUser(username string) ([]*Like, error)
}

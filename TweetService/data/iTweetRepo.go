package data

type TweetRepo interface {
	GetAll() Tweets
	CreateTweet(tw *Tweet)
	PutTweet(tw *Tweet, id int) error
	DeleteTweet(id int) error
}

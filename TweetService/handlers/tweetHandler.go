package handlers

import (
	"TweetService/data"
	"TweetService/proto/auth"
	"TweetService/proto/tweet"
	"context"
	"log"
)

type TweetHandler struct {
	tweet.UnimplementedTweetServiceServer
	l        *log.Logger
	repoImpl data.TweetRepo
	ac       auth.AuthServiceClient
}
type KeyProduct struct {
}

func NewTweet(l *log.Logger, repoImpl data.TweetRepo, ac auth.AuthServiceClient) *TweetHandler {
	return &TweetHandler{
		l:        l,
		repoImpl: repoImpl,
		ac:       ac,
	}
}

func (t *TweetHandler) GetTweets(ctx context.Context, r *tweet.GetTweetRequest) (*tweet.GetTweetResponse, error) {
	t.l.Println("Tweet service - Get tweets by username")
	// tweets := data.GetAll()

	tweets, err := t.repoImpl.GetTweetsByUsername(r.Username)
	if err != nil {
		t.l.Println("Error getting Tweets by username from cassandra")
		return nil, err

	}
	return &tweet.GetTweetResponse{
		TweetList: tweets,
	}, nil
}

func (t *TweetHandler) PostTweet(ctx context.Context, r *tweet.PostTweetRequest) (*tweet.PostTweetResponse, error) {
	t.l.Println("Tweet service - Post tweet")
	resp, err := t.ac.GetUser(context.Background(), &auth.UserRequest{
		Token: r.Token,
	})
	if err != nil {
		t.l.Println("Error getting user")
		return nil, err
	}
	res := &data.Tweet{
		Text:     r.Text,
		Username: resp.Username,
	}
	errorcic := t.repoImpl.CreateTweet(res)
	if errorcic != nil {
		t.l.Println("Error occurred during tweet creation")
		return nil, errorcic
	}
	return &tweet.PostTweetResponse{}, nil
	//data.RenderJson(w, tweet)
}

func (t *TweetHandler) LikeTweet(ctx context.Context, r *tweet.LikeTweetRequest) (*tweet.LikeTweetResponse, error) {
	t.l.Println("Tweet service - Like Tweet")
	resp, err := t.ac.GetUser(context.Background(), &auth.UserRequest{
		Token: r.Token,
	})
	if err != nil {
		t.l.Println("Error getting user")
		return nil, err
	}

	err1 := t.repoImpl.LikeTweet(r.TweetID, resp.Username, r.Like)

	if err1 != nil {
		t.l.Println("Error occurred during liking tweet")
		return nil, err1
	}

}

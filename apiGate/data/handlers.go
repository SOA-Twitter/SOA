package data

import (
	"TweeterMicro/TweetService/proto/tweet"
	"context"
	"log"
	"net/http"
)

type TweetHandler struct {
	l  *log.Logger
	pr tweet.TweetServiceClient
}
type TweetDAO struct {
	Text    string `json:"text"`
	Picture string `json:"picture"`
}

func NewTweetHandler(l *log.Logger, pr tweet.TweetServiceClient) *TweetHandler {
	return &TweetHandler{l, pr}
}
func (tw *TweetHandler) GetTweets(w http.ResponseWriter, r *http.Request) {
	resp, err := tw.pr.GetTweets(context.Background(), nil)
	if err != nil {
		tw.l.Println("Error getting tweets")
		http.Error(w, "Error getting tweets", http.StatusNotFound)
		return
	}
	tw.l.Println("Resp", resp)
}
func (tw *TweetHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	dao := TweetDAO{}
	err := FromJSON(&dao, r.Body)
	if err != nil {
		tw.l.Println("Cannot unmarshal json")
		http.Error(w, "Cannot unmarshal json", http.StatusBadRequest)
		return
	}
	_, err = tw.pr.PostTweet(context.Background(), &tweet.PostTweetRequest{
		Text:    dao.Text,
		Picture: dao.Picture,
	})
	if err != nil {
		tw.l.Println("Error occurred during creating tweet")
		return
	}
	w.Write([]byte("200 - CREATED"))

}

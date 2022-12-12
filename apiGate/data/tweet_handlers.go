package data

import (
	"apiGate/protos/tweet"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TweetHandler struct {
	l  *log.Logger
	pr tweet.TweetServiceClient
}

func NewTweetHandler(l *log.Logger, pr tweet.TweetServiceClient) *TweetHandler {
	return &TweetHandler{l, pr}
}
func (tw *TweetHandler) GetTweets(w http.ResponseWriter, r *http.Request) {
	tw.l.Println("Api-gate - Get tweets")
	username := mux.Vars(r)["username"]
	resp, err := tw.pr.GetTweets(context.Background(), &tweet.GetTweetRequest{
		Username: username,
	})
	if err != nil {
		tw.l.Println("Error getting tweets")
		http.Error(w, "Error getting tweets", http.StatusNotFound)
		return
	}

	err = ToJSON(resp.TweetList, w)
	tw.l.Println("Resp", resp)
}
func (tw *TweetHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	tw.l.Println("Api-gate - Create tweet")

	dao := Tweet{}
	err := FromJSON(&dao, r.Body)
	if err != nil {
		tw.l.Println("Cannot unmarshal json")
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	c := r.Header.Get("Authorization")
	if c == "" {
		http.Error(w, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
		return
	}

	_, err = tw.pr.PostTweet(context.Background(), &tweet.PostTweetRequest{
		Text:  dao.Text,
		Token: c,
	})
	if err != nil {
		tw.l.Println("Error occurred during creating tweet")
		return
	}
	w.Write([]byte("200 - CREATED"))

}

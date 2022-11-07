package handlers

import (
	"TweeterMicro/TweetService/data"
	"TweeterMicro/TweetService/proto/tweet"
	"context"
	"log"
	"net/http"
)

type TweetHandler struct {
	tweet.UnimplementedTweetServiceServer
	L        *log.Logger
	RepoImpl data.TweetRepo
}
type KeyProduct struct {
}

func NewTweet(l *log.Logger, repoImpl data.TweetRepo) *TweetHandler {
	return &TweetHandler{
		L:        l,
		RepoImpl: repoImpl}
}

func (t *TweetHandler) GetTweets(ctx context.Context, r *tweet.GetTweetRequest) (*tweet.GetTweetResponse, error) {
	t.L.Println("Handle GET tweet")
	// tweets := data.GetAll()

	tweets := t.RepoImpl.GetAll()
	return &tweet.GetTweetResponse{
		TweetList: tweets,
	}, nil
}
func (t *TweetHandler) PostTweet(ctx context.Context, r *tweet.PostTweetRequest) (*tweet.PostTweetResponse, error) {
	t.L.Println("Handle POST tweet")
	res := &data.Tweet{
		Text:    r.Text,
		Picture: r.Picture,
	}
	// data.CreateTweet(tweet)
	err := t.RepoImpl.CreateTweet(res)
	if err != nil {
		t.L.Println("Error occurred during tweet creation")
		return nil, err
	}
	return &tweet.PostTweetResponse{
		Id:     res.Id,
		Status: http.StatusCreated,
	}, nil
	//data.RenderJson(w, tweet)
}

func (p *TweetHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		tweet, err := data.DecodeBody(h.Body)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			p.L.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, tweet)
		h = h.WithContext(ctx)

		next.ServeHTTP(w, h)
	})
}
func (p *TweetHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.L.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

package handlers

import (
	"TweetService/data"
	"TweetService/proto/auth"
	"TweetService/proto/tweet"
	"context"
	"log"
	"net/http"
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
	t.l.Println("Handle GET tweet")
	// tweets := data.GetAll()

	tweets := t.repoImpl.GetAll()
	return &tweet.GetTweetResponse{
		TweetList: tweets,
	}, nil
}
func (t *TweetHandler) PostTweet(ctx context.Context, r *tweet.PostTweetRequest) (*tweet.PostTweetResponse, error) {
	t.l.Println("Handle POST tweet")
	resp, err := t.ac.GetUserId(context.Background(), &auth.UserIdRequest{
		Token: r.Token,
	})
	if err != nil {
		t.l.Println("Error creating tweet")
		t.l.Println(err)
		return nil, err
	}
	res := &data.Tweet{
		Text:      r.Text,
		Picture:   r.Picture,
		UserEmail: resp.UserEmail,
	}
	errorcic := t.repoImpl.CreateTweet(res)
	if errorcic != nil {
		t.l.Println("Error occurred during tweet creation")
		return nil, errorcic
	}
	return &tweet.PostTweetResponse{}, nil
	//data.RenderJson(w, tweet)
}

func (t *TweetHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		tweet, err := data.DecodeBody(h.Body)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			t.l.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, tweet)
		h = h.WithContext(ctx)

		next.ServeHTTP(w, h)
	})
}
func (t *TweetHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		t.l.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

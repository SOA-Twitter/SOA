package handlers

import (
	"TweeterMicro/TweetService/data"
	"context"
	"log"
	"net/http"
)

type TweetHandler struct {
	l        *log.Logger
	repoImpl data.TweetRepo
}
type KeyProduct struct {
}

func NewTweetHandler(l *log.Logger, twRepo data.TweetRepo) *TweetHandler {
	// *TODO test repoImpl Dependency Injection line behavior
	return &TweetHandler{l, twRepo}
}

func (t *TweetHandler) GetTweets(w http.ResponseWriter, h *http.Request) {
	t.l.Println("Handle GET tweet")
	// tweets := data.GetAll()
	tweets := t.repoImpl.GetAll()
	data.RenderJson(w, tweets)
}
func (t *TweetHandler) CreateTweet(w http.ResponseWriter, h *http.Request) {
	t.l.Println("Handle POST tweet")
	tweet := h.Context().Value(KeyProduct{}).(*data.Tweet)
	// data.CreateTweet(tweet)
	t.repoImpl.CreateTweet(tweet)
	w.WriteHeader(http.StatusCreated)
	//data.RenderJson(w, tweet)
}

func (p *TweetHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		tweet, err := data.DecodeBody(h.Body)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			p.l.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, tweet)
		h = h.WithContext(ctx)

		next.ServeHTTP(w, h)
	})
}
func (p *TweetHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.l.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

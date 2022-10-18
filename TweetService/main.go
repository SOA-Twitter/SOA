package main

import (
	"TweeterMicro/TweetService/handlers"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	l := log.New(os.Stdout, "Tweet-Api", log.LstdFlags)
	tweetHandler := handlers.NewTweetHandler(l)
	r := mux.NewRouter()
	r.Use(tweetHandler.MiddlewareContentTypeSet)
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/all", tweetHandler.GetTweets)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/createTweet", tweetHandler.CreateTweet)
	postRouter.Use(tweetHandler.MiddlewareProductValidation)

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

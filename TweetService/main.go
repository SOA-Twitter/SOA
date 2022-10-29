package main

import (
	"TweeterMicro/TweetService/data"
	"TweeterMicro/TweetService/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("app_port")
	if len(port) == 0 {
		port = "8080"
	}

	l := log.New(os.Stdout, "[Tweet-Api] ", log.LstdFlags)

	// *Dependency Injection of DB-communication into TweetHandler's repoImpl field
	// assign either NewPostgreSQL or NewInMemory to tweetRepoImpl
	tweetRepoImpl, err := data.CassandraConnection(l)
	if err != nil {
		l.Fatal(err)
	}
	tweetHandler := handlers.NewTweetHandler(l, &tweetRepoImpl)

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
	l.Println("Server listening on port ", port)

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

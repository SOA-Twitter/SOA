package main

import (
	"TweeterMicro/TweetService/proto/tweet"
	"TweeterMicro/apiGate/data"
	"context"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	l := log.New(os.Stdout, "[API_GATE] ", log.LstdFlags)

	tweetConn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		l.Println("ERooor")
	}
	defer tweetConn.Close()
	tweetClient := tweet.NewTweetServiceClient(tweetConn)
	tweetHandler := data.NewTweetHandler(l, tweetClient)
	r := mux.NewRouter()

	tweetRouter := r.PathPrefix("/tweet").Subrouter()
	tweetRouter.HandleFunc("/getTweets", tweetHandler.GetTweets).Methods(http.MethodGet)
	tweetRouter.HandleFunc("/postTweets", tweetHandler.PostTweet).Methods(http.MethodPost)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	l.Println("Server listening on port 8080")

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

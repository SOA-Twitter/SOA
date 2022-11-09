package main

import (
	"TweeterMicro/AuthService/proto/auth"
	"TweeterMicro/TweetService/proto/tweet"
	"TweeterMicro/apiGate/data"
	"context"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	l := log.New(os.Stdout, "[API_GATE] ", log.LstdFlags)

	authConn, err := grpc.DialContext(
		context.Background(),
		"localhost:8081",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatalf("Error connecting to Auth_Service: %v\n", err)
	}
	defer authConn.Close()
	authClient := auth.NewAuthServiceClient(authConn)
	authHandler := data.NewAuthHandler(l, authClient)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)

	//--------------------------------------------------------
	tweetConn, err := grpc.DialContext(
		context.Background(),
		"localhost:9092",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatalf("Error connecting to Tweet_Service: %v\n", err)
	}
	defer tweetConn.Close()
	tweetClient := tweet.NewTweetServiceClient(tweetConn)
	tweetHandler := data.NewTweetHandler(l, tweetClient)

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

package main

import (
	"apiGate/data"
	"apiGate/protos/auth"
	"apiGate/protos/tweet"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	port := os.Getenv("API_GATE_PORT")
	if len(port) == 0 {
		port = ":8082"
	}
	authHost := os.Getenv("AUTH_HOST")
	authPort := os.Getenv("AUTH_PORT")

	l := log.New(os.Stdout, "[API_GATE] ", log.LstdFlags)

	authConn, err := grpc.DialContext(
		context.Background(),
		authHost+":"+authPort,
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

	tweetPort := os.Getenv("TWEET_PORT")
	tweetHost := os.Getenv("TWEET_HOST")
	tweetConn, err := grpc.DialContext(
		context.Background(),
		tweetHost+":"+tweetPort,
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
	tweetRouter.Use(authHandler.VerifyJwt)
	tweetRouter.HandleFunc("/getTweets", tweetHandler.GetTweets).Methods(http.MethodGet)
	tweetRouter.HandleFunc("/postTweets", tweetHandler.PostTweet).Methods(http.MethodPost)

	s := &http.Server{
		Addr:         port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	l.Println("Server listening on port" + port)

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

package main

import (
	"TweeterMicro/TweetService/data"
	"TweeterMicro/TweetService/handlers"
	"TweeterMicro/TweetService/proto/tweet"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	port := os.Getenv("app_port")
	if len(port) == 0 {
		port = "8080"
	}
	l := log.New(os.Stdout, "[Tweet-Api] ", log.LstdFlags)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tweetRepoImpl, err := data.CassandraConnection(l)
	if err != nil {
		log.Println("Error connecting...")
	}
	tweetHandler := handlers.NewTweet(l, &tweetRepoImpl)
	tweet.RegisterTweetServiceServer(grpcServer, tweetHandler)
	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("server error: ", err)

		}
	}()
	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)
	<-stopCh
	grpcServer.Stop()

	//
	//port := os.Getenv("app_port")
	//if len(port) == 0 {
	//	port = "8080"
	//}
	//
	//l := log.New(os.Stdout, "[Tweet-Api] ", log.LstdFlags)
	//
	//// *Dependency Injection of DB-communication into TweetHandler's repoImpl field
	//// assign either NewPostgreSQL or NewInMemory to tweetRepoImpl
	//tweetRepoImpl, err := data.CassandraConnection(l)
	//if err != nil {
	//	l.Fatal(err)
	//}
	//tweetHandler := handlers.NewTweetHandler(l, &tweetRepoImpl)
	//
	//r := mux.NewRouter()
	//r.Use(tweetHandler.MiddlewareContentTypeSet)
	//getRouter := r.Methods(http.MethodGet).Subrouter()
	//getRouter.HandleFunc("/all", tweetHandler.GetTweets)
	//
	//postRouter := r.Methods(http.MethodPost).Subrouter()
	//postRouter.HandleFunc("/createTweet", tweetHandler.CreateTweet)
	//postRouter.Use(tweetHandler.MiddlewareProductValidation)

}

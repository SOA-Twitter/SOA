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

	grpcServer := grpc.NewServer()

	l := log.New(os.Stdout, "[Tweet-Api] ", log.LstdFlags)
	tweetRepoImpl, err := data.CassandraConnection(l)
	if err != nil {
		log.Println("Error connecting to cassandra...")
	}

	//PORT FIXED FOR NOW
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tweetHandler := handlers.NewTweet(l, &tweetRepoImpl)
	tweet.RegisterTweetServiceServer(grpcServer, tweetHandler)
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("Server error: ", err)

		}
	}()
	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)
	<-stopCh
	grpcServer.Stop()

	//postRouter.Use(tweetHandler.MiddlewareProductValidation)

}

package main

import (
	"TweetService/data"
	"TweetService/handlers"
	"TweetService/proto/auth"
	"TweetService/proto/tweet"
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

	port := os.Getenv("TWEET_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	l := log.New(os.Stdout, "[Tweet-Api] ", log.LstdFlags)
	tweetRepoImpl, err := data.CassandraConnection(l)
	if err != nil {
		l.Println("Error connecting to cassandra...")
	}
	authPort := os.Getenv("AUTH_PORT")

	conn, err := grpc.Dial(authPort+":"+authPort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ac := auth.NewAuthServiceClient(conn)

	//PORT FIXED FOR NOW
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tweetHandler := handlers.NewTweet(l, &tweetRepoImpl, ac)
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

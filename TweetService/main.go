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
	authHost := os.Getenv("AUTH_HOST")

	conn, err := grpc.Dial(authHost+":"+authPort, grpc.WithInsecure())
	if err != nil {
		l.Println("error connecting to auth service")
	}
	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	authHost+":"+authPort,
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	if err != nil {
		l.Println("error connecting to auth service")
		l.Println(err)
		l.Println(conn)
	}
	defer conn.Close()
	tweetRepoImpl.CreateTable()

	ac := auth.NewAuthServiceClient(conn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	tweetHandler := handlers.NewTweet(l, tweetRepoImpl, ac)
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

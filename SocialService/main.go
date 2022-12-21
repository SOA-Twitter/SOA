package main

import (
	"SocialService/data"
	"SocialService/handlers"
	"SocialService/proto/auth"
	"SocialService/proto/profile"
	"SocialService/proto/social"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("SOCIAL_PORT")
	if len(port) == 0 {
		port = "8083"
	}
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := log.New(os.Stdout, "[Social_Api] ", log.LstdFlags)
	socialRepoImpl, err := data.Neo4JConnection(l)
	if err != nil {
		l.Println("Error connecting to Neo4J")
	}
	defer socialRepoImpl.CloseDriverConnection(timeoutContext)
	socialRepoImpl.CheckConnection()
	////CONNECTION WITH AUTH SERVICE

	authPort := os.Getenv("AUTH_PORT")
	authHost := os.Getenv("AUTH_HOST")

	conn, err := grpc.Dial(authHost+":"+authPort, grpc.WithInsecure())
	if err != nil {
		l.Println("error connecting to auth service")
	}
	if err != nil {
		l.Println("error connecting to auth service")
		l.Println(err)
		l.Println(conn)
	}
	defer conn.Close()
	ac := auth.NewAuthServiceClient(conn)
	//-------------------------------------------------------------------

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	ps := profile.NewProfileServiceClient(conn)
	socialHandler := handlers.NewSocialHandler(l, socialRepoImpl, ac, ps)
	social.RegisterSocialServiceServer(grpcServer, socialHandler)
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

}

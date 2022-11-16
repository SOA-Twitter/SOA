package main

import (
	"ProfileService/data"
	"ProfileService/handlers"
	"ProfileService/proto/auth"
	"ProfileService/proto/profile"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("PROFILE_PORT")
	if len(port) == 0 {
		port = "9099"
	}
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	l := log.New(os.Stdout, "[Profile-Api]", log.LstdFlags)
	profileRepoImpl, err := data.MongoConnection(timeoutContext, l)
	if err != nil {
		l.Println("Error connecting to Mongo...")
	}
	defer profileRepoImpl.Disconnect(timeoutContext)

	//----------Connection with Auth Service----------
	authPort := os.Getenv("AUTH_PORT")
	authHost := os.Getenv("AUTH_HOST")

	conn, err := grpc.DialContext(
		context.Background(),
		authHost+":"+authPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Println("error connecting to auth service")
		l.Println(err)
	}
	defer conn.Close()
	as := auth.NewAuthServiceClient(conn)
	//-------------------------------------------------

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		l.Fatalf("Failed to listen")
	}
	grpcServer := grpc.NewServer()
	profileHandler := handlers.NewProfileHandler(l, profileRepoImpl, as)
	profile.RegisterProfileServiceServer(grpcServer, profileHandler)
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

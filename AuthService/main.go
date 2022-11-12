package main

import (
	"AuthService/data"
	"AuthService/handlers"
	"AuthService/proto/auth"
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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8081"
	}
	grpcServer := grpc.NewServer()
	l := log.New(os.Stdout, "[Auth-Api]", log.LstdFlags)
	authRepo, err := data.PostgresConnection(l)
	if err != nil {
		l.Println("Error connecting to postgres...")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}
	authHandler := handlers.NewAuthHandler(l, authRepo)
	auth.RegisterAuthServiceServer(grpcServer, authHandler)
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

package main

import (
	"AuthService/data"
	"AuthService/handlers"
	"AuthService/proto/auth"
	"AuthService/proto/profile"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := os.Getenv("AUTH_PORT")
	if len(port) == 0 {
		port = "8002"
	}
	grpcServer := grpc.NewServer()
	l := log.New(os.Stdout, "[Auth-Api]", log.LstdFlags)
	authRepo, err := data.PostgresConnection(l)
	if err != nil {
		l.Println("Error connecting to postgres...")
	}
	profilePort := os.Getenv("PROFILE_PORT")
	profileHost := os.Getenv("PROFILE_HOST")

	conn, err := grpc.Dial(profileHost+":"+profilePort, grpc.WithInsecure())
	if err != nil {
		l.Println("error connecting to profile service")
	}
	defer conn.Close()
	ps := profile.NewProfileServiceClient(conn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}
	authHandler := handlers.NewAuthHandler(l, authRepo, ps)
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

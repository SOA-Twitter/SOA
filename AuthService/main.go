package main

import (
	"TweeterMicro/AuthService/data"
	"TweeterMicro/AuthService/handlers"
	"TweeterMicro/AuthService/proto/auth"
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

	port := os.Getenv("Auth_port")
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
	authHandler := handlers.NewAuthHandler(l, &authRepo)
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

	//l := log.New(os.Stdout, "[Auth-Api]", log.LstdFlags)
	//authRepo, err := data.PostgresConnection(l)
	//if err != nil {
	//	log.Println("Error connecting to postgres...")
	//}
	//authHandler := handlers.NewAuthHandler(l, &authRepo)
	//
	//r := mux.NewRouter()
	//s := r.Methods(http.MethodPost).Subrouter()
	//s.HandleFunc("/login", authHandler.Login)
	//s.HandleFunc("/register", authHandler.Register)
	//
	//m := r.Methods(http.MethodGet).Subrouter()
	//m.HandleFunc("/home", handlers.Home)
	//m.Use(middleware.VerifyJwt)
	//srv := &http.Server{
	//	Addr:         ":" + port,
	//	Handler:      r,
	//	IdleTimeout:  120 * time.Second,
	//	ReadTimeout:  1 * time.Second,
	//	WriteTimeout: 1 * time.Second,
	//}
	//l.Println("Server listening on port", port)
	//
	//go func() {
	//	err := srv.ListenAndServe()
	//	if err != nil {
	//		l.Fatal(err)
	//	}
	//}()
	//sigChan := make(chan os.Signal)
	//signal.Notify(sigChan, syscall.SIGINT)
	//signal.Notify(sigChan, syscall.SIGTERM)
	//
	//sig := <-sigChan
	//l.Println("Graceful shutdown", sig)
	//
	//tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//srv.Shutdown(tc)

}

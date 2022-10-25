package main

import (
	"TweeterMicro/auth/handlers"
	"TweeterMicro/auth/middleware"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	port := os.Getenv("Auth_port")
	if len(port) == 0 {
		port = "8081"
	}
	l := log.New(os.Stdout, "[Auth-Api]", log.LstdFlags)
	authHandler := handlers.NewAuthHandler(l)

	r := mux.NewRouter()
	s := r.Methods(http.MethodPost).Subrouter()
	s.HandleFunc("/login", authHandler.Login)
	s.HandleFunc("/register", authHandler.Register)

	m := r.Methods(http.MethodGet).Subrouter()
	m.HandleFunc("/home", handlers.Home)
	m.Use(middleware.VerifyJwt)
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
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
	srv.Shutdown(tc)

}

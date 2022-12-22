package main

import (
	"apiGate/data"
	"apiGate/protos/auth"
	"apiGate/protos/profile"
	"apiGate/protos/social"
	"apiGate/protos/tweet"
	"context"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	port := os.Getenv("API_GATE_PORT")
	if len(port) == 0 {
		port = ":8082"
	}
	authHost := os.Getenv("AUTH_HOST")
	authPort := os.Getenv("AUTH_PORT")

	l := log.New(os.Stdout, "[API_GATE] ", log.LstdFlags)

	authConn, err := grpc.DialContext(
		context.Background(),
		authHost+":"+authPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatalf("Error connecting to Auth_Service: %v\n", err)
	}
	defer authConn.Close()
	authClient := auth.NewAuthServiceClient(authConn)
	authHandler := data.NewAuthHandler(l, authClient)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)
	authRouter.HandleFunc("/activate/{activationId}", authHandler.ActivateProfile).Methods(http.MethodGet)
	authRouter.HandleFunc("/recoverEmail", authHandler.SendRecoveryEmail).Methods(http.MethodPost)
	authRouter.HandleFunc("/recover", authHandler.RecoverProfile).Methods(http.MethodPost)
	//--------------------------------------------------------

	tweetPort := os.Getenv("TWEET_PORT")
	tweetHost := os.Getenv("TWEET_HOST")
	tweetConn, err := grpc.DialContext(
		context.Background(),
		tweetHost+":"+tweetPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatalf("Error connecting to Tweet_Service: %v\n", err)
	}
	defer tweetConn.Close()
	tweetClient := tweet.NewTweetServiceClient(tweetConn)
	tweetHandler := data.NewTweetHandler(l, tweetClient)

	tweetRouter := r.PathPrefix("/tweet").Subrouter()
	tweetRouter.Use(authHandler.Authorize)
	tweetRouter.HandleFunc("/getTweets/{username}", tweetHandler.GetTweetsByUsername).Methods(http.MethodGet)
	tweetRouter.HandleFunc("/getTweetLikes/{id}", tweetHandler.GetLikesByTweetId).Methods(http.MethodGet)
	tweetRouter.HandleFunc("/postTweets", tweetHandler.PostTweet).Methods(http.MethodPost)
	tweetRouter.HandleFunc("/like/{id}", tweetHandler.LikeTweet).Methods(http.MethodPut)

	//----------------------------------------------------------SOCIAL SERVICE

	socialPort := os.Getenv("SOCIAL_PORT")
	socialHost := os.Getenv("SOCIAL_HOST")
	socialConn, err := grpc.DialContext(
		context.Background(),
		socialHost+":"+socialPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatalf("Error connecting to Social_Service: %v\n", err)
	}
	defer socialConn.Close()
	socialClient := social.NewSocialServiceClient(socialConn)
	socialHandler := data.NewSocialHandler(l, socialClient)

	socialRouter := r.PathPrefix("/social").Subrouter()
	tweetRouter.Use(authHandler.Authorize)
	socialRouter.HandleFunc("/follow", socialHandler.Follow).Methods(http.MethodPost)
	socialRouter.HandleFunc("/unfollow", socialHandler.Unfollow).Methods(http.MethodDelete)
	socialRouter.HandleFunc("/pending", socialHandler.GetPendingFollowRequests).Methods(http.MethodGet)

	defer socialConn.Close()

	//----------------------------------------------------------
	profileHost := os.Getenv("PROFILE_HOST")
	profilePort := os.Getenv("PROFILE_PORT")
	profileConn, err := grpc.DialContext(
		context.Background(),
		profileHost+":"+profilePort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatalf("Error connecting to Profile_service: %v\n", err)
	}
	defer profileConn.Close()

	profileClient := profile.NewProfileServiceClient(profileConn)
	profileHandler := data.NewProfileHandler(l, profileClient)

	profileRouter := r.PathPrefix("/profile").Subrouter()
	profileRouter.Use(authHandler.Authorize)
	profileRouter.HandleFunc("/{username}", profileHandler.UserProfile).Methods(http.MethodGet)
	profileRouter.HandleFunc("/changePassword", authHandler.ChangePassword).Methods(http.MethodPost)
	profileRouter.HandleFunc("/privacy", profileHandler.ManagePrivacy).Methods(http.MethodPut)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"https://localhost:4200"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		gorillaHandlers.AllowCredentials())

	s := &http.Server{
		Addr:         port,
		Handler:      cors(r),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
	}
	l.Println("Server listening on port" + port)

	go func() {
		err := s.ListenAndServeTLS("certificates/cert.crt", "certificates/cert.key")
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
	s.Shutdown(tc)

}

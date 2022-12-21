package data

import (
	"apiGate/protos/social"
	"context"
	"log"
	"net/http"
)

type SocialHandler struct {
	l  *log.Logger
	pr social.SocialServiceClient
}

func NewSocialHandler(l *log.Logger, pr social.SocialServiceClient) *SocialHandler {
	return &SocialHandler{
		l, pr}
}

func (h *SocialHandler) Follow(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("Social - Follow Handler")

	c := request.Header.Get("Authorization")
	if c == "" {
		http.Error(writer, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
		return
	}

	userToFollow := UserNode{}
	err := FromJSON(&userToFollow, request.Body)
	if err != nil {
		h.l.Println(unMarshall)
		http.Error(writer, invalidJson, http.StatusBadRequest)
		return
	}

	response, errSocial := h.pr.RequestToFollow(context.Background(), &social.FollowRequest{
		Username: userToFollow.Username,
		Token:    c,
	})
	if errSocial != nil {
		h.l.Println("Unable to Follow that user: ")

		if response.Status == "" {
			h.l.Println("Unavailable for following - user could not be found")
			http.Error(writer, "User could not be found", http.StatusBadRequest)
		} else {
			h.l.Println("Unavailable for following - user could not be found")
			http.Error(writer, "You've already requested to follow this user: Status "+response.Status, http.StatusNotAcceptable)
		}

		return
	}

	// TODO look how "response" JSON is shown on client side !
	err = ToJSON(response, writer)
	if err != nil {
		h.l.Println(jsonErrMsg)
		http.Error(writer, jsonErrMsg, http.StatusInternalServerError)
		return
	}
}

func (h *SocialHandler) Unfollow(writer http.ResponseWriter, request *http.Request) {
	h.l.Println("Social - Unfollow Handler")

	c := request.Header.Get("Authorization")
	if c == "" {
		http.Error(writer, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
		return
	}

	userToUnfollow := UserNode{}
	err := FromJSON(&userToUnfollow, request.Body)
	if err != nil {
		h.l.Println(unMarshall)
		http.Error(writer, invalidJson, http.StatusBadRequest)
		return
	}

	_, errSocial := h.pr.Unfollow(context.Background(), &social.FollowRequest{
		Username: userToUnfollow.Username,
		Token:    c,
	})
	if errSocial != nil {
		h.l.Println("Unable to Unfollow that user: ")
		http.Error(writer, "Cannot unfollow that user", http.StatusNotFound)
		return
	}
}

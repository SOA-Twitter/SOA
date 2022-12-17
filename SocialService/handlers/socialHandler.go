package handlers

import (
	"SocialService/data"
	"SocialService/proto/auth"
	"SocialService/proto/social"
	"context"
	"log"
)

type SocialHandler struct {
	social.UnimplementedSocialServiceServer
	l        *log.Logger
	repoImpl data.SocialRepo
	ac       auth.AuthServiceClient
}

func NewSocialHandler(l *log.Logger, repoImpl data.SocialRepo, ac auth.AuthServiceClient) *SocialHandler {
	return &SocialHandler{
		l:        l,
		repoImpl: repoImpl,
		ac:       ac,
	}
}
func (s *SocialHandler) AddUser(ctx context.Context, r *social.RegUserRequest) (*social.RegUserResponse, error) {
	s.l.Println("Social service - Add user node to neo4J")

	//TODO

	return &social.RegUserResponse{}, nil
}

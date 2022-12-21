package handlers

import (
	"SocialService/data"
	"SocialService/proto/auth"
	"SocialService/proto/profile"
	"SocialService/proto/social"
	"context"
	"log"
)

type SocialHandler struct {
	social.UnimplementedSocialServiceServer
	l        *log.Logger
	repoImpl data.SocialRepo
	ac       auth.AuthServiceClient
	pc       profile.ProfileServiceClient
}

func NewSocialHandler(l *log.Logger, repoImpl data.SocialRepo, ac auth.AuthServiceClient, pc profile.ProfileServiceClient) *SocialHandler {
	return &SocialHandler{
		l:        l,
		repoImpl: repoImpl,
		ac:       ac,
		pc:       pc,
	}
}

func (s *SocialHandler) RegUser(ctx context.Context, r *social.RegUserRequest) (*social.RegUserResponse, error) {
	s.l.Println("Social service - Register User")

	err := s.repoImpl.RegUser(r.Username)
	if err != nil {
		return &social.RegUserResponse{}, err
	}
	return &social.RegUserResponse{}, nil
}

func (s *SocialHandler) RequestToFollow(ctx context.Context, r *social.FollowRequest) (*social.FollowIntentResponse, error) {
	s.l.Println("Social service - Follow user intent")

	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		s.l.Println("Error getting claims")
		return &social.FollowIntentResponse{}, err
	}

	foundUser, err := s.pc.GetUserProfile(context.Background(), &profile.UserProfRequest{
		Username: r.Username,
	})
	if err != nil {
		s.l.Println("Social handler - Error fetching user from Profile Service: ", err)
		return &social.FollowIntentResponse{}, err
	}

	followReqStatus, err := s.repoImpl.Follow(claims.Username, r.Username, foundUser.Private)
	if err != nil {
		return &social.FollowIntentResponse{
			Status: followReqStatus,
		}, err
	}
	return &social.FollowIntentResponse{
		Status: followReqStatus,
	}, nil
}

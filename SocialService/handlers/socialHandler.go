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

	err = s.repoImpl.Follow(claims.Username, r.Username, foundUser.Private)
	if err != nil {
		return &social.FollowIntentResponse{
			//Status: followReqStatus,
		}, err
	}
	return &social.FollowIntentResponse{
		//Status: followReqStatus,
	}, nil
}

func (s *SocialHandler) Unfollow(ctx context.Context, r *social.FollowRequest) (*social.UnfollowResponse, error) {
	s.l.Println("Social service - Unfollow user")

	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		s.l.Println("Error getting claims")
		return &social.UnfollowResponse{}, err
	}

	err1 := s.repoImpl.Unfollow(claims.Username, r.Username)
	if err1 != nil {
		s.l.Println("Cannot unfollow ", r.Username)
		return &social.UnfollowResponse{}, err1
	}
	return &social.UnfollowResponse{}, nil
}

func (s *SocialHandler) GetPendingFollowRequests(ctx context.Context, r *social.GetPendingRequest) (*social.PendingFollowerResponse, error) {
	s.l.Println("Social service - Get pending follow requests")

	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		s.l.Println("Error getting claims")
		return &social.PendingFollowerResponse{}, err
	}

	result, err := s.repoImpl.GetPendingFollowers(claims.Username)
	if err != nil {
		s.l.Println("Cannot Get Pending Followers")
		return &social.PendingFollowerResponse{}, err
	}
	return &social.PendingFollowerResponse{
		PendingFollowers: result,
	}, nil
}

func (s *SocialHandler) IsFollowed(ctx context.Context, r *social.IsFollowedRequest) (*social.IsFollowedResponse, error) {
	s.l.Println("Social service - Is Followed by logged user")

	s.l.Println("Social Svc - JWToken: ", r.Requester)
	claims, err := data.GetFromClaims(r.Requester)
	if err != nil {
		s.l.Println("Error getting claims")
		return nil, err
	}

	s.l.Println("Social Svc - Logged: ", claims.Username, "Target user: ", r.Target)
	result, err := s.repoImpl.IsFollowed(claims.Username, r.Target)
	if err != nil {
		s.l.Println("Error returning Is Followed info")
		return nil, err
	}
	return &social.IsFollowedResponse{
		IsFollowedByLogged: result,
	}, nil
}

func (s *SocialHandler) DeclineFollowRequest(ctx context.Context, r *social.ManageFollowRequest) (*social.ManageFollowResponse, error) {
	s.l.Println("Social service - Decline Follow Request")

	claims, err := data.GetFromClaims(r.Requester)
	if err != nil {
		s.l.Println("Error getting claims")
		return &social.ManageFollowResponse{}, err
	}

	err1 := s.repoImpl.DeclineFollowRequest(claims.Username, r.Target)
	if err1 != nil {
		s.l.Println("Error Declining follow request from ", claims.Username)
		return &social.ManageFollowResponse{}, err1
	}
	return &social.ManageFollowResponse{}, nil
}

func (s *SocialHandler) AcceptFollowRequest(ctx context.Context, r *social.ManageFollowRequest) (*social.ManageFollowResponse, error) {
	s.l.Println("Social service - Accept Follow Request")

	claims, err := data.GetFromClaims(r.Requester)
	if err != nil {
		s.l.Println("Error getting claims")
		return &social.ManageFollowResponse{}, err
	}

	err1 := s.repoImpl.AcceptFollowRequest(claims.Username, r.Target)
	if err1 != nil {
		s.l.Println("Error Accepting follow request from  ", claims.Username)
		return &social.ManageFollowResponse{}, err1
	}
	return &social.ManageFollowResponse{}, nil
}

func (s *SocialHandler) HomeFeed(ctx context.Context, r *social.HomeFeedRequest) (*social.HomeFeedResponse, error) {
	s.l.Println("Social service - Home feed")
	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		s.l.Println("Error getting claims")
		return &social.HomeFeedResponse{}, err
	}

	usernames, err1 := s.repoImpl.GetFollowers(claims.Username)
	if err1 != nil {
		s.l.Println("Error getting followers from database")
		return &social.HomeFeedResponse{}, err1
	}

	return &social.HomeFeedResponse{
		Usernames: usernames,
	}, nil

}

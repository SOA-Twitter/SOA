package handlers

import (
	"TweeterMicro/ProfileService/data"
	"TweeterMicro/ProfileService/proto/profile"
	"context"
	"log"
)

type ProfileHandler struct {
	profile.UnimplementedProfileServiceServer
	l    *log.Logger
	repo data.ProfileRepo
}

func NewProfileHandler(l *log.Logger, repo data.ProfileRepo) *ProfileHandler {
	return &ProfileHandler{
		l:    l,
		repo: repo,
	}
}
func (pr *ProfileHandler) Register(ctx context.Context, r *profile.RegisterRequest) (*profile.RegisterResponse, error) {
	pr.l.Println("Register handler")

	//TODO REGISTER HANDLER

	return &profile.RegisterResponse{}, nil
}

func (pr *ProfileHandler) GetUserProfile(ctx context.Context, r *profile.UserProfRequest) (*profile.UserProfResponse, error) {
	pr.l.Println("Register handler")

	//TODO USERPROFILE HANDLER

	return &profile.UserProfResponse{}, nil

}

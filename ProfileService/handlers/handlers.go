package handlers

import (
	"ProfileService/data"
	"ProfileService/proto/auth"
	"ProfileService/proto/profile"
	"context"
	"log"
)

type ProfileHandler struct {
	profile.UnimplementedProfileServiceServer
	l    *log.Logger
	repo *data.ProfileRepo
	as   auth.AuthServiceClient
}

func NewProfileHandler(l *log.Logger, repo *data.ProfileRepo, as auth.AuthServiceClient) *ProfileHandler {
	return &ProfileHandler{
		l:    l,
		repo: repo,
		as:   as,
	}
}
func (pr *ProfileHandler) Register(ctx context.Context, r *profile.ProfileRegisterRequest) (*profile.ProfileRegisterResponse, error) {
	pr.l.Println("Register handler")
	user := &data.User{
		Username:       r.Username,
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		Email:          r.Email,
		Gender:         r.Gender,
		Country:        r.Country,
		Age:            int(r.Age),
		CompanyName:    r.CompanyName,
		CompanyWebsite: r.CompanyWebsite,
		Private:        r.Private,
	}
	err := pr.repo.Register(user)
	if err != nil {
		pr.l.Println("Error inserting user into db")
		return nil, err
	}

	return &profile.ProfileRegisterResponse{}, nil
}

func (pr *ProfileHandler) GetUserProfile(ctx context.Context, r *profile.UserProfRequest) (*profile.UserProfResponse, error) {
	pr.l.Println("Get User profile handler")
	user, err := pr.repo.GetByUsername(r.Username)
	if err != nil {
		pr.l.Println("Cannot find user")
		return nil, err
	}
	return &profile.UserProfResponse{
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Gender:         user.Gender,
		Country:        user.Country,
		Age:            int32(user.Age),
		CompanyName:    user.CompanyName,
		CompanyWebsite: user.CompanyWebsite,
		Private:        user.Private,
	}, nil

}

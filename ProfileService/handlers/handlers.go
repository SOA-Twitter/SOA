package handlers

import (
	"ProfileService/data"
	"ProfileService/proto/auth"
	"ProfileService/proto/profile"
	social "ProfileService/proto/social"
	"context"
	"log"
	"net/http"
)

type ProfileHandler struct {
	profile.UnimplementedProfileServiceServer
	l    *log.Logger
	repo *data.ProfileRepo
	as   auth.AuthServiceClient
	ss   social.SocialServiceClient
}

func NewProfileHandler(l *log.Logger, repo *data.ProfileRepo, as auth.AuthServiceClient, ss social.SocialServiceClient) *ProfileHandler {
	return &ProfileHandler{
		l:    l,
		repo: repo,
		as:   as,
		ss:   ss,
	}
}
func (pr *ProfileHandler) Register(ctx context.Context, r *profile.ProfileRegisterRequest) (*profile.ProfileRegisterResponse, error) {
	pr.l.Println("Profile service - Register")
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
		Private:        true,
		Role:           r.Role,
	}
	err := pr.repo.Register(user)
	if err != nil {
		pr.l.Println("Error inserting user into db")
		return nil, err
	}

	_, err6 := pr.ss.RegUser(context.Background(), &social.RegUserRequest{
		Username: r.Username,
	})
	if err6 != nil {
		return nil, err6
	}

	return &profile.ProfileRegisterResponse{}, nil
}

func (pr *ProfileHandler) GetUserProfile(ctx context.Context, r *profile.UserProfRequest) (*profile.UserProfResponse, error) {
	pr.l.Println("Profile service - Get User Profile")

	user, err1 := pr.repo.GetByUsername(r.Username)
	if err1 != nil {
		pr.l.Println("Cannot find user")
		return nil, err1
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
		Role:           user.Role,
	}, nil

}

func (pr *ProfileHandler) ManagePrivacy(ctx context.Context, r *profile.ManagePrivacyRequest) (*profile.ManagePrivacyResponse, error) {
	pr.l.Println("Profile service - Manage privacy")
	claims, err := data.GetFromClaims(r.Token)
	if err != nil {
		pr.l.Println("Error getting claims")
		return &profile.ManagePrivacyResponse{
			Status: http.StatusNotFound,
		}, err
	}

	err1 := pr.repo.ChangePrivacy(claims.Username, r.Privacy)

	if err1 != nil {
		pr.l.Println("Cannot update account privacy for not-found user " + claims.Username)
		return nil, err1
	}
	return &profile.ManagePrivacyResponse{
		Privacy: r.Privacy,
		Status:  http.StatusOK,
	}, nil
}

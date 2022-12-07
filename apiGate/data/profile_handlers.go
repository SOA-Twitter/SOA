package data

import (
	"apiGate/protos/profile"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ProfileHandler struct {
	l  *log.Logger
	pr profile.ProfileServiceClient
}

func NewProfileHandler(l *log.Logger, pr profile.ProfileServiceClient) *ProfileHandler {
	return &ProfileHandler{l, pr}
}

func (ah *ProfileHandler) UserProfile(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("Api-gate - Get User Info")
	username := mux.Vars(r)["username"]
	response, err := ah.pr.GetUserProfile(context.Background(), &profile.UserProfRequest{
		Username: username,
	})
	if err != nil {
		ah.l.Println("Error occurred. Cannot get user infos!")
		return
	}
	userProfile := UserInfo{
		Username:       response.Username,
		FirstName:      response.FirstName,
		LastName:       response.LastName,
		Email:          response.Email,
		Gender:         Gender(response.Gender),
		Country:        response.Country,
		Age:            int(response.Age),
		CompanyName:    response.CompanyName,
		CompanyWebsite: response.CompanyWebsite,
		Private:        response.Private,
	}
	err = ToJSON(userProfile, w)
	if err != nil {
		ah.l.Println("Unable to convert to json")
		http.Error(w, "Unable to convert to json", http.StatusInternalServerError)
		return
	}
}

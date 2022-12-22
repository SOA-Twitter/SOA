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

const (
	jsonErrMsg = "Unable to convert to json"
)

func NewProfileHandler(l *log.Logger, pr profile.ProfileServiceClient) *ProfileHandler {
	return &ProfileHandler{l, pr}
}

func (ah *ProfileHandler) UserProfile(w http.ResponseWriter, r *http.Request) {
	ah.l.Println("Api-gate - User Profile")
	username := mux.Vars(r)["username"]

	c := r.Header.Get("Authorization")
	if c == "" {
		http.Error(w, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
		return
	}

	response, err := ah.pr.GetUserProfile(context.Background(), &profile.UserProfRequest{
		Username: username,
		Token:    c,
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
		ah.l.Println(jsonErrMsg)
		http.Error(w, jsonErrMsg, http.StatusInternalServerError)
		return
	}
}

func (ah *ProfileHandler) ManagePrivacy(w http.ResponseWriter, req *http.Request) {
	ah.l.Println("Api-gate - Manage Privacy")

	privacy := ManagePrivacy{}

	err1 := FromJSON(&privacy, req.Body)
	if err1 != nil {
		ah.l.Println(jsonErrMsg)
		http.Error(w, jsonErrMsg, http.StatusInternalServerError)
		return
	}

	c := req.Header.Get("Authorization")
	if c == "" {
		http.Error(w, "Unauthorized! NO COOKIE", http.StatusUnauthorized)
		return
	}
	response, err2 := ah.pr.ManagePrivacy(context.Background(), &profile.ManagePrivacyRequest{
		Token:   c,
		Privacy: privacy.Private,
	})

	if err2 != nil {
		ah.l.Println("Cannot change account privacy!")
		http.Error(w, "Cannot change your account privacy!", http.StatusInternalServerError)
		return
	}

	err3 := ToJSON(response, w)
	if err3 != nil {
		ah.l.Println(jsonErrMsg)
		http.Error(w, jsonErrMsg, http.StatusInternalServerError)
		return
	}

}

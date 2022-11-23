package data

import (
	"apiGate/protos/profile"
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

func (ah *ProfileHandler) Register1(w http.ResponseWriter, r *http.Request) {

}

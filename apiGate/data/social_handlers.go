package data

import (
	"apiGate/protos/social"
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

func (h SocialHandler) Follow(writer http.ResponseWriter, request *http.Request) {

}

package data

import "google.golang.org/protobuf/types/known/timestamppb"

type Tweet struct {
	Id           string                `json:"id"`
	Text         string                `json:"text" validate:"required"`
	Username     string                `json:"username"`
	CreationDate timestamppb.Timestamp `json:"creationDate"`
}

type Like struct {
	TweetId  string `json:"tweetId"`
	Username string `json:"username"`
	Liked    bool   `json:"liked"`
}

type TokenStr struct {
	Token string `json:"token"`
}

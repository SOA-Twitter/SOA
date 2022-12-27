package data

import "SocialService/proto/social"

type SocialRepo interface {
	RegUser(username string) error
	Follow(usernameOfFollower string, usernameToFollow string, isPrivate bool) error
	Unfollow(usernameOfRequester string, usernameToUnfollow string) error
	GetPendingFollowers(usernameOfRequester string) ([]*social.PendingFollower, error)
	IsFollowed(requesterUsername string, targetUsername string) (bool, error)
	DeclineFollowRequest(usernameOfFollowed string, usernameOfFollower string) error
	AcceptFollowRequest(usernameOfFollowed string, usernameOfFollower string) error
}

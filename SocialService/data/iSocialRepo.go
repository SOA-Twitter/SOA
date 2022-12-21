package data

type SocialRepo interface {
	RegUser(username string) error
	Follow(usernameOfFollower string, usernameToFollow string, isPrivate bool) (string, error)
	Unfollow(usernameOfRequester string, usernameToUnfollow string) error
}
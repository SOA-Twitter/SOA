package data

type Tweet struct {
	Id           string `json:"id"`
	Text         string `json:"text" validate:"required"`
	Username     string `json:"username"`
	CreationDate string `json:"creationdate"`
}

type Like struct {
	TweetId  string `json:"tweetId"`
	Username string `json:"username"`
	Liked    bool   `json:"liked"`
}

type TokenStr struct {
	Token string `json:"token"`
}

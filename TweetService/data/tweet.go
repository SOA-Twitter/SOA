package data

type Tweet struct {
	ID        int    `json:"id"`
	Text      string `json:"text" validate:"required"`
	Picture   string `json:"picture"`
	CreatedOn string `json:"-"`
	DeletedOn string `json:"deleted_on"`
}

package models

type User struct {
	ID        int    `json:"id"`
	TagName   string `json:"tag_name"`
	AvatarURL string `json:"avatar_url"`
}

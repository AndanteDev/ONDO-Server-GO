package model

import "time"

type DBHandler interface {
}
type MysqlHandler struct {
}
type user struct {
	Id       int    `json:"id"`
	Mail     string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Access_token  string    `json:"access_token"`
	Token_type    string    `json:"token_type"`
	Expiration    time.Time `json:"expires_in"`
	Refresh_token string    `json:"refresh_token"`
}
type OauthConfig struct {
	Code          string
	Client_id     string
	Client_secret string
	Redirect_uri  string
	Grant_type    string
}

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
	AccessToken  string    `json:"access_token" binding:"required"`
	TokenType    string    `json:"token_type" binding:"required"`
	Expiration   time.Time `json:"expires_in" binding:"required"`
	RefreshToken string    `json:"refresh_token" binding:"required"`
}
type OauthConfig struct {
	Code          string
	Client_id     string
	Client_secret string
	Scope         []string
	Redirect_uri  string
	Grant_type    string
}

package handler

import (
	"ondo/server/go/info"
	"ondo/server/go/utils"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googlewebOauthConfig = oauth2.Config{
	RedirectURL:  info.GoogleRedirectPath,
	ClientID:     os.Getenv("webclient_id"),
	ClientSecret: os.Getenv("webclient_secret"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func webOauth(c *gin.Context) string {

	accessToken, refreshToken := utils.GenerateStateOauthToken(rdb, c)
	url := googlewebOauthConfig.AuthCodeURL(state)
	return url
}

func andOauth(c *gin.Context) {

}
func iosOauth(c *gin.Context) {

}

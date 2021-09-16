package handler

import (
	"log"
	"net/http"
	"ondo/server/go/info"
	"ondo/server/go/utils"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthConn interface {
	loginHandler(c *gin.Context)
}
type Oauth struct {
	oauth OauthConn
	os    string
}

func kakaoLoginHandler(c *gin.Context) {

}
func googleLoginHandler(c *gin.Context) {
	var googleOauthConfig = oauth2.Config{
		RedirectURL:  info.GoogleRedirectPath,
		ClientID:     os.Getenv("webclient_id"),
		ClientSecret: os.Getenv("webclient_secret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	state := utils.GenerateStateOauthToken
	url := googleOauthConfig.AuthCodeURL()

}

func googleCallBackHandler(c *gin.Context) {
	c.Redirect(http.StatusOK, "/")
}
func kakaoCallBackHandler(c *gin.Context) {

}

func indexhandler(c *gin.Context) {
	log.Println(c.Request.Header.Get("User-Agent"))
}

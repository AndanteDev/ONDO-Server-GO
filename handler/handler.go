package handler

import (
	"net/http"
	"ondo/server/go/info"
	"ondo/server/go/model"
	"ondo/server/go/utils"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &model.OauthConfig{
	Client_id:     os.Getenv("webclient_id"),
	Client_secret: os.Getenv("webclient_secret"),
	Scope:         []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Redirect_uri:  info.GoogleRedirectPath,
	Grant_type:    "authorization_code",
}

var googleTestOauthConfig = &oauth2.Config{
	RedirectURL:  info.GoogleRedirectPath,
	ClientID:     os.Getenv("webclient_id"),
	ClientSecret: os.Getenv("webclient_secret"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func testgoogleLoginHandler(c *gin.Context) {
	state := utils.GenerateStateOauthCookie(c)
	url := googleTestOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleLoginHandler(c *gin.Context) {

}

//receive code and get accesstoken
func googleCallBackHandler(c *gin.Context) {
	googleUserInfo, err := utils.GetGoogleUserInfo(googleTestOauthConfig, c)
	utils.HandleErr(http.StatusNotFound, err, c)
	c.JSON(http.StatusOK, string(googleUserInfo))
}

func kakaoCallBackHandler(c *gin.Context) {

}

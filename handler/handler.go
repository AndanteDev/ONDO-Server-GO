package handler

import (
	"fmt"
	"log"
	"net/http"
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

func kakaoLoginHandler(c *gin.Context) {

}
func googleLoginHandler(c *gin.Context) {
	url := webOauth(c)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func webOauth(c *gin.Context) string {
	state := utils.GenerateOauthState(c, rdb)
	url := googlewebOauthConfig.AuthCodeURL(state)
	fmt.Println(url)
	return url
}

func googleCallBackHandler(c *gin.Context) {
	oauthstate := utils.GetOauthState(c, rdb)
	if c.Request.FormValue("state") != oauthstate {
		log.Printf("invaild google state Token:%s state:%s", oauthstate, c.Request.FormValue("state"))
		c.Redirect(http.StatusFound, "/")
		return
	}
	data, err := utils.GetGoogleUserInfo(c, c.Request.FormValue("code"), googlewebOauthConfig)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.JSON(http.StatusOK, data)
}

func kakaoCallBackHandler(c *gin.Context) {

}

func indexHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/index.html")
}

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
	ClientSecret: "8NwJMGbg30V9Z8QwXo26z8cY",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func kakaoLoginHandler(c *gin.Context) {

}
func googleLoginHandler(c *gin.Context) {
	state := utils.GenerateOauthState(c, rdb)
	url := googlewebOauthConfig.AuthCodeURL(state)
	fmt.Println(url)
	c.Redirect(http.StatusFound, url)
}

func googleCallBackHandler(c *gin.Context) {
	oauthstate := utils.GetOauthState(c, rdb)
	fmt.Println(oauthstate, c.Request.FormValue("state"))
	if c.Request.FormValue("state") != oauthstate {
		log.Printf("invaild google state Token:%s state:%s", oauthstate, c.Request.FormValue("state"))
		c.Redirect(http.StatusFound, "/")
		return
	}
	data, err := utils.GetGoogleUserInfo(c, c.Request.FormValue("code"), &googlewebOauthConfig)
	fmt.Println(c.Request.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/")
		return
	}
	fmt.Fprintln(c.Writer, string(data))
	c.JSON(http.StatusOK, data)
}

func kakaoCallBackHandler(c *gin.Context) {

}

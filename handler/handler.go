package handler

import (
	"log"
	"net/http"
	"ondo/server/go/httputil"
	"ondo/server/go/info"
	"ondo/server/go/model"
	"ondo/server/go/utils"
	"os"

	"github.com/gin-gonic/gin"
)

var googleOauthConfig = &model.OauthConfig{
	Client_id:     os.Getenv("webclient_id"),
	Client_secret: os.Getenv("webclient_secret"),
	Scope:         []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Redirect_uri:  info.GoogleRedirectPath,
	Grant_type:    "authorization_code",
}

//receive code and get accesstoken
func googleCallBackHandler(c *gin.Context) {
	googleOauthConfig.Code = c.Query("code")
	googletoken, err := utils.GetGoogleAccessToken(googleOauthConfig)
	if err != nil {
		httputil.NewRedirect()
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/")
		return
	}
	jwtstring, err := utils.GetGoogleUserInfoJWT(googletoken)

	if err != nil {
		httputil.NewRedirect()
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.JSON(http.StatusOK, map[string]string{"jwtstring": jwtstring})

}

func kakaoCallBackHandler(c *gin.Context) {

}

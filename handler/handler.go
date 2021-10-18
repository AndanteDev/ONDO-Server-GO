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
	Client_secret: "",
	Redirect_uri:  info.GoogleRedirectPath,
	Grant_type:    "authorization_code",
}

func googleCallBackHandler(c *gin.Context) {
	googleOauthConfig.Code = c.Query("code")
	token, err := utils.GetGoogleAccessToken(googleOauthConfig)
	if err != nil {

		httputil.NewRedirect()
		log.Println(err.Error())
		c.Redirect(http.StatusFound, "/")
		return
	}
	utils.GetGoogleUserInfo(token)
}

func kakaoCallBackHandler(c *gin.Context) {

}

package handler

import (
	"log"
	"net/http"
	"ondo/server/go/utils"

	"github.com/gin-gonic/gin"
)

func kakaoLoginHandler(c *gin.Context) {

}
func googleLoginHandler(c *gin.Context) {
	os := c.Request.Header.Get("User-Agent")
	utils.IdentifyOS(os)
	url := webOauth(c)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleCallBackHandler(c *gin.Context) {
	oauthstate := rdb.Get(c, "oauthstate")
	if c.Request.FormValue("state") != oauthstate.Val() {
		log.Printf("invaild google state Token:%s state:%s", oauthstate.Val(), c.Request.FormValue("state"))
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	data, err := utils.GetGoogleUserInfo(c, c.Request.FormValue("code"), googlewebOauthConfig)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	c.JSON(http.StatusOK, data)
}

func kakaoCallBackHandler(c *gin.Context) {

}

func indexhandler(c *gin.Context) {
	log.Println(c.Request.Header.Get("User-Agent"))
}

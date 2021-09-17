package handler

import (
	"log"
	"net/http"
	"ondo/server/go/utils"

	"github.com/gin-gonic/gin"
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
	os := c.Request.Header.Get("User-Agent")
	utils.IdentifyOS(os)
	url := webOauth(c)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleCallBackHandler(c *gin.Context) {
	oauthstate := rdb.Get(c, "oauthstate")
	if c.Request.FormValue("state") != oauthstate.Val() {
		log.Println("invaild google state Token:", oauthstate.Val())
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}
func kakaoCallBackHandler(c *gin.Context) {

}

func indexhandler(c *gin.Context) {
	log.Println(c.Request.Header.Get("User-Agent"))
}

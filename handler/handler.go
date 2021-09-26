package handler

import (
	"log"
	"net/http"

type OauthConn interface {
	loginHandler(c *gin.Context)
}
type Oauth struct {
	oauth OauthConn
	os    string
}
func googleLoginHandler(c *gin.Context) {
	os := c.Request.Header.Get("User-Agent")
	utils.IdentifyOS(os)
	url := webOauth(c)
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
func googleLoginHandler(c *gin.Context) {


func kakaoCallBackHandler(c *gin.Context) {

}

func indexhandler(c *gin.Context) {
	log.Println(c.Request.Header.Get("User-Agent"))
}

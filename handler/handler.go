package handler

import (
	"log"
	"net/http"

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

}

func googleCallBackHandler(c *gin.Context) {
	c.Redirect(http.StatusOK, "/")
}
func kakaoCallBackHandler(c *gin.Context) {

}

func indexhandler(c *gin.Context) {
	log.Println(c.Request.Header.Get("User-Agent"))
}

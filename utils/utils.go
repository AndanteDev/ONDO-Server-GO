package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"ondo/server/go/info"
	"ondo/server/go/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GetOauthState(c *gin.Context, rdb *redis.Client) string {
	state := rdb.Get(c, "oauthstate").Val()

	return state
}
func GenerateOauthState(c *gin.Context, rdb *redis.Client) string {
	expiration := 30 * time.Minute
	state := randValue()
	rdb.Set(c, "oauthstate", state, expiration).Val()

	return state
}
func IdentifyOS(os string) string {
	return ""
}

func randValue() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

func GetGoogleAccessToken(ac *model.OauthConfig) (*model.Token, error) {
	pbytes, _ := json.Marshal(ac)
	fmt.Println(string(pbytes))
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(info.GetGoogleTokenURL, "application/json", buff)
	if err != nil {
		return nil, errors.New(err.Error() + " Get Token Error")
	}

	token := &model.Token{}
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error() + " Token ResponseBody error")
	}
	err = json.Unmarshal(respbody, token)
	if err != nil {
		return nil, errors.New(err.Error() + " Token Unmarshal error")
	}
	return token, nil
}

func GetGoogleUserInfoJWT(token *model.Token) (string, error) {
	resp, err := http.Get(info.GetGoogleUserURL + token.AccessToken)
	if err != nil {
		return nil, errors.New(err.Error() + " Get UserInfo Error")
	}

	respbody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.New(err.Error() + " Userinfo ResponseBody error")
	}
	token, err := createUserToken(respbody)
	return respbody, nil
}

func createUserToken(respbody []byte) (token string, err error) {

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

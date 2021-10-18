package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"ondo/server/go/info"
	"ondo/server/go/model"
	"time"

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
	token := &model.Token{}

	pbytes, _ := json.Marshal(token)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(info.GetGoogleTokenURL, "application/json", buff)
	if err != nil {
		return nil, err
	}
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respbody, token)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetGoogleUserInfo(token *model.Token) ([]byte, error) {
	info.GetGoogleUserURL
	return nil, nil
}

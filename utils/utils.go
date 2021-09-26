package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/oauth2"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func GetOauthState(c *gin.Context) string {
	state := rdb.Get(c, "oauthstate").Val()

	return state
}
func GenerateOauthState(c *gin.Context) string {
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

func GetGoogleUserInfo(c *gin.Context, code string, ac oauth2.Config) ([]byte, error) {
	token, err := ac.Exchange(c, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s", err.Error())
	}

}

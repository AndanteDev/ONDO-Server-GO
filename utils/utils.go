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

func GenerateStateOauthToken(rdb *redis.Client, c *gin.Context) (string, string) {
	accexpiration := 30 * time.Minute
	accessToken := randValue()

	refexpiration := 60 * 24 * time.Hour
	refreshToken := randValue()

	rdb.Set(c, "oauthstate_accessToken", accessToken, accexpiration).Val()
	rdb.Set(c, "oauthstate_refreshToken", refreshToken, refexpiration).Val()

	return accessToken, refreshToken
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

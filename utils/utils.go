package utils

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GenerateStateOauthToken(rdb *redis.Client, c *gin.Context) string {
	expiration := 1 * 24 * time.Hour
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

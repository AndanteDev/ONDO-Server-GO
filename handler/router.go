package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func MakeHandler() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.Group("/auth")
		{
			google := v1.Group("/google")
			{
				google.GET("callback", googleCallBackHandler)
			}
			kakao := v1.Group("/kakao")
			{
				kakao.GET("callback", kakaoCallBackHandler)
			}
		}

	}
	return r
}

package handler

import "github.com/gin-gonic/gin"

func MakeHandler() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.Group("/auth")
		{
			v1.Group("/google")
			{
				v1.GET("/login", googleLoginHandler)
				v1.GET("callback", googleCallBackHandler)
			}
			v1.Group("/kakao")
			{
				v1.GET("/login", kakaoLoginHandler)
				v1.GET("callback", kakaoCallBackHandler)

			}
		}

	}

	return r
}

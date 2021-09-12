package handler

import "github.com/gin-gonic/gin"

func MakeHandler() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/signin", signinHandler)
		v1.POST("/signup", signupHandler)
	}

	return r
}

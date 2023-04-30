package router

import (
	"imgo/api"
	"imgo/middleware"

	"github.com/gin-gonic/gin"
)

// 路由配置
func Init() *gin.Engine {
	r := gin.Default()

	r.POST("/register", api.UserRegister)
	r.POST("login", api.UserLogin)
	r.GET("/ws", api.WSServer)
	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		v1.POST("/request", api.PushRequest)
		v1.PUT("/accept", api.AcceptRequest)
		v1.POST("/group", api.CreateGroup)
		v1.POST("/group_member", api.JoinGroup)

	}

	return r
}

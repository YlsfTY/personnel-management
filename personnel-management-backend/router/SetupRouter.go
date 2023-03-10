package router

import (
	"personnel-management-backend/controller"
	"personnel-management-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (r *gin.Engine) {
	// 创建一个默认引擎，有Logger(日志) 与 Recovery(崩溃处理)中间件
	r = gin.Default()
	api := r.Group("/api") //路由组
	{
		api.GET("/ping", controller.Ping)
		api.POST("/user/register", controller.Register) //注册
		api.POST("/user/login", controller.Login)       //登录
		api.GET("/user/info", middleware.AuthMiddleware(), controller.Info)
	}
	return
}

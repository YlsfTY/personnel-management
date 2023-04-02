package router

import (
	"fmt"
	"net/http"
	"personnel-management-backend/controller"
	"personnel-management-backend/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// 创建一个默认引擎，有Logger(日志) 与 Recovery(崩溃处理)中间件
	r.Use(middleware.CorsMiddleware())
	api := r.Group("/api") //路由组
	{
		api.GET("/ping", controller.Ping)
		api.POST("/user/register", controller.UserRegister) //注册
		api.POST("/user/login", controller.UserLogin)       //登录
		api.POST("/user/test", func(ctx *gin.Context) {
			code := ctx.PostForm("code")
			count, _ := strconv.Atoi(code)
			fmt.Println(count)
			// ctx.JSON(count, gin.H{
			// 	"code": code,
			// })
			ctx.AbortWithStatus(http.StatusNotFound)
		})
		// api.GET("/user/info", middleware.AuthMiddleware(), controller.Info)
	}
	{
		api.Use(middleware.AuthMiddleware())
		api.GET("/user/info", controller.Info)
	}
}

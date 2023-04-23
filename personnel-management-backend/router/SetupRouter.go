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
			ctx.AbortWithStatus(http.StatusNotFound)
		})
	}
	{
		api.Use(middleware.AuthMiddleware())
		api.GET("/user/info", controller.Info)
		api.POST("/personnel/create", controller.CreatePer)
		api.GET("/personnel/getPersonnel", controller.GetPer)
		api.POST("/personnel/uploadImage", controller.UploadImage)
		api.GET("/personnel/getImage", controller.GetImage)
	}
	static := r.Group("/static")
	{
		static.Use(func(c *gin.Context) {
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Header("Expires", "0")
			c.Header("Pragma", "no-cache")
			c.Next()
		})
		static.Static("/image", "./static/image")
	}
}

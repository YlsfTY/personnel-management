package controller

import (
	// "log"
	"net/http"
	"personnel-admin/dao"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("userName")
	password := ctx.PostForm("userPassword")
	user := dao.User{
		Name:     name,
		Password: password,
		State:    true,
	}

	// 数据验证
	if len(name) == 0 || len(password) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Account or password resolution failed",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Password cannot be less than 6 digits",
		})
		return
	}

	if dao.SearchUserName(name) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Duplicate usernames",
		})
		return
	}

	// log.Panicln(name,password)//日志打印

	// 创建用户
	dao.CreateUser(&user)
	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Register success!",
	})
}

func Login(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("userName")
	password := ctx.PostForm("userPassword")
	user := dao.User{
		Name:     name,
		Password: password,
		State:    true,
	}

	// 数据验证
	if len(name) == 0 || len(password) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Account or password resolution failed",
		})
		return
	}

	sign, _ := dao.CheckUser(&user)
	if !sign {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Wrong username or password",
		})
	}
	// 返回数据
	ctx.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":"Login success",
	})

}

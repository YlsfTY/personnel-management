package controller

import (
	// "log"
	"log"
	"net/http"
	"personnel-management-backend/common"
	"personnel-management-backend/dao"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("userName")
	password := ctx.PostForm("userPassword")
	user := &dao.User{
		Name:     name,
		Password: password,
		State:    true,
	}

	// 数据验证
	if len(name) == 0 || len(password) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
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

	if dao.SearchUserName(user) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Duplicate usernames",
		})
		return
	}
	// 密码加密（bcrypt加密）
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Encryption error",
		})
		return
	}
	user.Password = string(hasedPassword)

	// 创建用户
	_, err = dao.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to add user",
		})
		return
	}
	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Register success!",
	})
}

func Login(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("userName")
	password := ctx.PostForm("userPassword")
	user := &dao.User{
		Name:     name,
		Password: password,
	}

	// 数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "Account or password resolution failed",
		})
		return
	}
	// 判断用户是否存在
	sign := dao.SearchUserName(user)
	if !sign {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The user does not exist ",
		})
		return
	}
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "wrong password",
		})
		return
	}
	// 分析是否异常
	sign, _ = dao.CheckState(user)
	if !sign {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "User exception",
		})
		return
	}
	// 发放token
	token, err := common.CreateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "token err",
		})
		log.Printf("token generate error:%v", err)
		return
	}
	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"msg": "Login success",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": user,
		},
	})
}

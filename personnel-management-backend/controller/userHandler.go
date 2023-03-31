package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"personnel-management-backend/common"
	"personnel-management-backend/dao"
	"personnel-management-backend/ulits"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("userName")
	password := ctx.PostForm("userPassword")

	errMsg, ok := validateUserParams(name, password)
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, errMsg)
		return
	}

	// 密码加密（bcrypt加密）
	hasedPassword, err := hashPassword(password)
	if err != nil {
		errMsg = "Encryption error"
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, errMsg)
		return
	}
	user := &dao.User{
		Name:     name,
		Password: hasedPassword,
		State:    true,
	}

	// 创建用户
	ok, err = user.CreateUser()
	if err != nil {
		// 创建失败
		errMsg = "Failed to create user"
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, errMsg)
		return
	}
	if !ok {
		// 用户已经存在
		errMsg = "The user already exists."
		ulits.ResponseWithError(ctx, http.StatusConflict, errMsg)
		return
	}

	// 返回数据
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: http.StatusOK,
		Data: nil,
		Msg:  "Register success!",
	})
}

func UserLogin(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("userName")
	password := ctx.PostForm("userPassword")

	// 数据验证
	errMsg, ok := validateUserParams(name, password)
	if !ok {
		ulits.ResponseWithError(ctx, http.StatusBadRequest, errMsg)
		return
	}

	user := &dao.User{
		Name: name,
	}

	//  检验用户
	ok, err := user.CheckUserState()
	if err != nil {
		if errors.Is(err, fmt.Errorf("The user does not exist.")) {
			errMsg = err.Error()
			ulits.ResponseWithError(ctx, http.StatusConflict, errMsg)
			return
		}
		errMsg = "Internal Server Error."
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, errMsg)
		return
	} else if !ok {
		errMsg = "This account has been deleted."
		ulits.ResponseWithError(ctx, http.StatusForbidden, errMsg)
		return
	}

	// 验证密码
	if !verifPassword(user.Password, password) {
		errMsg = "wrong password"
		ulits.ResponseWithError(ctx, http.StatusBadRequest, errMsg)
		return
	}
	// 发放token
	token, err := common.CreateToken(user)
	if err != nil {
		errMsg = "Error generating token"
		log.Printf("token generate error:%v", err)
		ulits.ResponseWithError(ctx, http.StatusInternalServerError, errMsg)
		return
	}
	// 返回数据
	ulits.ResponseWithData(ctx, http.StatusOK, ulits.ResponseData{
		Code: 200,
		Data: gin.H{
			"token": token,
		},
		Msg: "Login success",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": user.(dao.User).Name,
		},
	})
}

func validateUserParams(name, password string) (string, bool) {
	if len(name) < 1 || len(name) > 16 {
		return "Account resolution failed", false
	} else if len(password) < 6 || len(password) > 16 {
		return "Password resolution failed", false
	}
	return "", true
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func verifPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

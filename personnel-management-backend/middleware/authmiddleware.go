package middleware

import (
	"errors"
	"log"
	"net/http"
	"personnel-management-backend/common"
	"personnel-management-backend/dao"
	"personnel-management-backend/ulits"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ulits.ResponseWithError(ctx, http.StatusUnauthorized, "权限不足")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		claims, err := common.ParseToken(tokenString)
		if err != nil {
			log.Printf("解析token失败:%s", err.Error())
			ulits.ResponseWithError(ctx, http.StatusUnauthorized, "解析token失败")
			ctx.Abort()
			return
		}
		userId := claims.UserId
		var user dao.User
		err = dao.Db.First(&user, userId).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("用户 %d 不存在", claims.UserId)
			} else {
				log.Printf("查询用户 %d 失败：%s", claims.UserId, err.Error())
			}
			ulits.ResponseWithError(ctx, http.StatusUnauthorized, "权限不足,")
			ctx.Abort()
			return
		}
		// 用户存在，信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}

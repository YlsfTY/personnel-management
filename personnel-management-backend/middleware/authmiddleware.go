package middleware

import (
	"net/http"
	"personnel-management-backend/common"
	"personnel-management-backend/dao"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		claims, err := common.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			ctx.Abort()
			return
		}
		userId := claims.UserId
		var user dao.User
		dao.MysqlDb.First(&user, userId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		// 用户存在，信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}

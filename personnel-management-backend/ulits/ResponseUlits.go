package ulits

import "github.com/gin-gonic/gin"

func ResponseWithError(ctx *gin.Context, statusCode int, errMsg string) {
	ResponseWithData(ctx, statusCode, ResponseData{
		Code: statusCode,
		Data: nil,
		Msg:  errMsg,
	})
}

func ResponseWithData(ctx *gin.Context, statusCode int, data ResponseData) {
	ctx.JSON(statusCode, data)
}

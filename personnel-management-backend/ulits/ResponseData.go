package ulits

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code int    `json:"code"`
	Data gin.H  `json:"data"`
	Msg  string `json:"msg"`
}

package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res.ReturnError(code))
}

// 通常成功数据处理
func OK(c *gin.Context, data interface{}) {
	var res Response
	res.Data = data
	res.Msg = ""
	c.JSON(http.StatusOK, res.ReturnOK())
}

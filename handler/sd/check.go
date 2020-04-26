package sd

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// Pesponse http 返回状态码
type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func(res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

// Health 健康检查
func HealthCheck(c *gin.Context) {
	res := Response{ Msg: "ok"}
	c.JSON(http.StatusOK, res.ReturnOK())
}
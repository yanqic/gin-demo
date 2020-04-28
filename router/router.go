package router

import (
	"gin-demo/controller/user"
	"gin-demo/middleware/jwt"
	"github.com/gin-gonic/gin"
	"log"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", user.Login)
	v1 := r.Group("/api/v1")
	v1.Use(jwt.Auth())
	{
		// 获取标签列表
		v1.GET("/user", user.GetUsers)
		v1.POST("/user", user.AddUser)
		v1.PUT("/user/:id", user.EditUser)
	}
	log.Println("路由加载成功！")
	return r
}

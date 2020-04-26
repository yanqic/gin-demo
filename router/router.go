package router

import (
	"gin-demo/handler/sd"
	"gin-demo/handler/user"
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

	// 监控信息
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		// svcd.GET("/disk", sd.DiskCheck)
		// svcd.GET("/cpu", sd.CPUCheck)
		// svcd.GET("/ram", sd.RAMCheck)
		// svcd.GET("/os", sd.OSCheck)
	}

	//r.POST("/login", user.Login)
	v1 := r.Group("/api/v1")
	// v1.Use(auth.JWTAuth())
	{
		// 获取标签列表
		v1.GET("/user", user.GetUser)
		//v1.POST("user", user.AddUser)
	}
	log.Println("路由加载成功！")
	return r
}
package user

import (
	"github.com/gin-gonic/gin"
)

// AddUser 添加用户 
func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"addUser": "succ",
	})
}

// GetUser 添加用户 
func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"getUser": "succ",
	})
}
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) (skip int, pageSize int) {
	pageSize = com.StrTo(c.Query("pageSize")).MustInt()
	if pageSize == 0 {
		pageSize = 20
	}
	skip = 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		skip = (page - 1) * pageSize
	}

	return skip, pageSize
}

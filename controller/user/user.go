package user

import (
	"gin-demo/config"
	MyJwt "gin-demo/middleware/jwt"
	"gin-demo/pkg/app"
	"gin-demo/pkg/util"
	UserService "gin-demo/service/user"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"time"
)

type userStruct struct {
	Id        int    `json:"id" binding:"omitempty"`
	Username  string `json:"username" binding:"required,min=1,max=20"`
	Password  string `json:"password" binding:"required,min=6,max=50"`
	CreatedBy int    `json:"created_by" binding:"omitempty"`
}

func Login(c *gin.Context) {
	var reqInfo userStruct
	if err := c.ShouldBindJSON(&reqInfo); err != nil {
		app.Error(c, http.StatusBadRequest, err, "")
		return
	}

	userService := UserService.User{Username: reqInfo.Username, Password: reqInfo.Password}

	userInfo, err := userService.Check()

	if err != nil {
		app.Error(c, http.StatusInternalServerError, err, "")
		return
	}

	j := &MyJwt.JWT{SigningKey: []byte(config.JwtConfig.Secret)}

	token, err := j.GenerateToken(userInfo.ID, userInfo.Username, userInfo.Password, time.Now())

	if err != nil {
		app.Error(c, http.StatusInternalServerError, err, "")
		return
	}

	app.OK(c, map[string]string{
		"token": token,
	})
}

// AddUser 添加用户
func AddUser(c *gin.Context) {

	var reqInfo userStruct
	if err := c.ShouldBindJSON(&reqInfo); err != nil {
		app.Error(c, http.StatusBadRequest, err, "请求体格式错误")
		return
	}
	claims := c.MustGet("claims").(*MyJwt.Claims)
	userService := UserService.User{
		Username:  reqInfo.Username,
		Password:  reqInfo.Password,
		CreatedBy: claims.Id,
	}
	if userInfo, err := userService.Add(); err == nil {

		app.OK(c, userInfo)
	} else {
		app.Error(c, http.StatusInternalServerError, err, "")
		return
	}
}

// GetUser 获取所有用户
func GetUsers(c *gin.Context) {
	skip, limit := util.GetPage(c)
	userService := UserService.User{
		ID:       com.StrTo(c.Query("id")).MustInt(),
		PageNum:  skip,
		PageSize: limit,
	}

	total, err := userService.Count()
	if err != nil {
		app.Error(c, http.StatusInternalServerError, err, "获取数量错误")
		return
	}

	userList, err := userService.GetAll()
	if err != nil {
		app.Error(c, http.StatusInternalServerError, err, "获取数据错误")
		return
	}
	for _, v := range userList {
		v.Password = ""
	}

	data := make(map[string]interface{})
	data["lists"] = userList
	data["total"] = total

	app.OK(c, data)
}

// EditUser 更新用户信息
func EditUser(c *gin.Context) {
	var reqInfo userStruct
	id := com.StrTo(c.Param("id")).MustInt()
	if err := c.ShouldBindJSON(&reqInfo); err != nil {
		app.Error(c, http.StatusBadRequest, err, "")
		return
	} else if id == 0 {
		app.Error(c, http.StatusBadRequest, err, "请求ID不存在")
		return
	}

	userService := UserService.User{
		ID:        id,
		Username:  reqInfo.Username,
		Password:  reqInfo.Password,
		CreatedBy: reqInfo.CreatedBy,
	}

	err := userService.Edit()

	if err != nil {
		app.Error(c, http.StatusInternalServerError, err, "")
		return
	}

	app.OK(c, map[string]string{"result": "ok"})
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func loginHandler(c *gin.Context) {
	// 获取用户POST过来的数据
	//username := c.PostForm("username")
	//password := c.DefaultPostForm("password", "xxx")

	// 利用gin框架提供的数据绑定 从请求中取出数据
	var u User
	err := c.ShouldBind(&u)
	if err != nil {
		c.String(http.StatusOK, "参数错误")
		return
	}
	// 参数校验方式：1.自己校验 2.通过第三方库在反序列化的时候校验
	if len(u.Username) == 0 {
		// 请求中没有传递username参数
		c.String(http.StatusOK, "username不能为空")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf(
		"username:%v password:%v",
		u.Username, u.Password))
}

func main() {
	r := gin.Default()
	r.POST("/login", loginHandler)
	r.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func postLoginHandler(c *gin.Context) {
	// 获取用户输入
	var u user
	err := c.ShouldBind(&u)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"errMessage": "请求参数有误！",
		})
		return
	}
	if len(u.Username) == 0 || len(u.Password) == 0 {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"errMessage": "用户名或密码为空",
		})
		return
	}
	// 去数据库校验一下用户名或密码是否正确

	// 登录成功跳转到index页面
	c.Redirect(http.StatusFound, "/index")
}

func main() {
	r := gin.Default()
	// 1. 加载静态文件
	r.Static("/assets", "./static")
	// 2. 自定义模板函数 （没有就省略这一步）
	// 3. 加载模板文件
	r.LoadHTMLGlob("templates/*")
	// GET:跟server端要一个登录的页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", postLoginHandler)
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run()
}

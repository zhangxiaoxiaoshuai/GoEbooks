package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1:before c.Next()")
	//c.Next() // 是中间件中请求和响应的分界点
	fmt.Println("m1:after c.Next()")
}

func m2(c *gin.Context) {
	fmt.Println("m2:before c.Next()")
	// 判断用户是否登录成功
	// 如果登录不成功就不让请求继续往后走
	//c.String(http.StatusOK, "请先登录！")
	//return
	//c.Abort() // 终止调用链条
	//c.AbortWithStatus(200)
	c.Set("username", "小p")
	fmt.Println("m2:after c.Next()")
}

// 利用闭包实现的中间件，注意注册的时候的写法
func xxx(ok bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ok {
			// 登录校验
			fmt.Println("进行登录校验")
		} else {
			fmt.Println("不用登录校验")
		}
	}
}

func indexHandler(c *gin.Context) {
	fmt.Println("index handler")
	username, ok := c.Get("username")
	if !ok {
		username = "匿名用户"
	}
	//username.(string) 是空接口需要进行类型转换
	c.String(http.StatusOK, fmt.Sprintf("hello %v", username))
}
func homeHandler(c *gin.Context) {
	fmt.Println("home handler")
	c.String(http.StatusOK, "home")
}

func main() {
	//gin.SetMode(gin.ReleaseMode) // 线上发布模式
	r := gin.Default()
	//r.Use(m1, m2) // 注册全局的中间件
	// 局部注册
	r.GET("/index", m1, m2, indexHandler)
	//r.GET("/home", m2, homeHandler)
	//// 路由组注册
	////shopGroup := r.Group("/shop", m1, m2)
	//shopGroup := r.Group("/shop")
	//shopGroup.Use(xxx(false))
	//{
	//	shopGroup.GET("/index", m1, m2, func(c *gin.Context) {
	//		c.String(http.StatusOK, "/shop/index")
	//	})
	//}

	r.Run() // 正常阻塞在这个地方

	//r2 := gin.Default()
	//r2.Run(":8081")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newsHandler(c *gin.Context) {
	yearValue := c.Param("year") // string
	// QueryString参数
	pageValue := c.Query("page")
	pagesizeValue := c.DefaultQuery("pagesize", "20")

	c.String(http.StatusOK, fmt.Sprintf(
		"year:%v page:%v pagesize:%v",
		yearValue,
		pageValue,
		pagesizeValue))
}

func loginHandler(c *gin.Context) {
	// 获取用户POST过来的数据
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "xxx")
	c.String(http.StatusOK, fmt.Sprintf(
		"username:%v password:%v",
		username, password))
}

func main() {
	r := gin.Default()
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", newsHandler)
		shopGroup.GET("/cart", newsHandler)
		shopGroup.GET("/order", newsHandler)
		shopGroup.GET("/rank", newsHandler)
		// 路由组继续拆分路由组
		loop2 := shopGroup.Group("/xx")
		loop2.GET("/loop2", newsHandler)
	}

	r.GET("/shop/index", newsHandler)
	r.GET("/shop/cart", newsHandler)
	r.GET("/shop/order", newsHandler)
	r.GET("/shop/rank", newsHandler)

	r.GET("/book/index", newsHandler)
	r.GET("/book/cart", newsHandler)
	r.GET("/book/order", newsHandler)
	r.GET("/book/rank", newsHandler)

	r.GET("/api/v1/index", newsHandler)
	r.GET("/api/v1/shop/cart", newsHandler)
	r.GET("/api/v1/shop/order", newsHandler)
	r.GET("/api/v1/shop/rank", newsHandler)

	r.POST("/login", loginHandler)
	r.Run()
}

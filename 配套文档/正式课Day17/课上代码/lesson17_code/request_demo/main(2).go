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
	r.GET("/news/:year", newsHandler)
	r.POST("/login", loginHandler)
	r.Run()
}

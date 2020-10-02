package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "xiaop",
		"age":18,
	})
}

func loginHandler(c *gin.Context){
	//c.Request.Method
}

func main() {
	router := gin.Default() // 生成默认的路由器
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/json", jsonHandler)

	router.Any("/login", loginHandler)

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})
	router.Run()
}
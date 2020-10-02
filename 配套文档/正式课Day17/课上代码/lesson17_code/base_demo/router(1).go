package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func someGetHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}

func SetupRoute() *gin.Engine {
	router := gin.Default() // 1. 比下面New()多了两个中间件，Logger()和Recovery()
	//router := gin.New()  // 2.

	router.GET("/someGet", someGetHandler)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	//
	//router.Any("/", handler)
	//
	//router.NoRoute(handler404)

	// 添加 Get 请求路由
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	return router
}

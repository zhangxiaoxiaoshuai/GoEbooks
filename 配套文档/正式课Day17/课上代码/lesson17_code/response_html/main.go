package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 在加载模板文件之前去注册自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//加载模板
	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/login.html", "templates/index.html")
	//定义路由
	shopGroup := router.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			//根据完整文件名渲染模板，并传递参数
			c.HTML(http.StatusOK, "shop/index.html", gin.H{
				"title": "Main website",
			})
		})
	}
	router.Run(":8080")
}

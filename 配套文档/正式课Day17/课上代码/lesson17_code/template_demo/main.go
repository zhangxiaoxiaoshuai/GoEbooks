package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/multitemplate"
)

// loadTemplates 自定义加载模板文件的函数
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layout/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/include/*.html")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func main() {
	r := gin.Default()
	r.Static("/xxx", "./static")
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)
	r.Run()
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		// 上传文件 发送POST请求来了
		// 1. 获取用户上传的文件
		fh, _ := c.FormFile("xx")
		fmt.Println(fh.Filename, fh.Size)

		// 2. server端保存
		err := c.SaveUploadedFile(fh, fh.Filename)
		if err != nil {
			c.String(http.StatusOK, "上传失败")
			return
		}
		// 3. 返回响应
		c.String(http.StatusOK, "上传成功")
	} else {
		// 想要访问上传文件的页面
		c.HTML(http.StatusOK, "upload.html", nil)
	}
}

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("upload.html")

	r.Any("/upload", uploadHandler)

	r.Run()
}

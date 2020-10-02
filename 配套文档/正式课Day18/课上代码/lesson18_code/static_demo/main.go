package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "index")
}

func main() {
	r := gin.Default()
	r.Static("/static", "/dist")
	//r.StaticFile("/static")
	r.StaticFS("/showDir", http.Dir("."))
	r.GET("/", indexHandler)
	r.Run()
}

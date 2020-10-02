package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// 定义一个自定义的函数
func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func main() {
	// 1. 将自定义的函数注册到模板引擎中
	funcMap := template.FuncMap {
		"fdate": formatDate,
	}
	// 2. 解析模板
	t := template.New("test.tpl").Funcs(funcMap)
	t, err := t.ParseFiles("test.tpl")
	if err != nil {
		log.Fatal("Parse errr:", err)
	}

	err = t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal("Exeute error:", err)
	}
}
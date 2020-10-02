package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func stringLiteralTemplate() {
	// 1. 定义模板
	s := "My name is {{ .Name }}. I am {{ .Age }} years old.\n"
	// 2. 解析模板
	t, err := template.New("test").Parse(s)
	if err != nil {
		log.Fatal("Parse string literal template error:", err)
	}
	// 3.渲染模板
	u := User{Name: "lianshi", Age: 18}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute string literal template error:", err)
	}
}

func fileTemplate() {
	t, err := template.ParseFiles("test.html")
	if err != nil {
		log.Fatal("Parse file template error:", err)
	}

	u := User{Name: "ls", Age: 18}
	err = t.Execute(os.Stdout, u)
	if err != nil {
		log.Fatal("Execute file template error:", err)
	}
}

func main() {
	//stringLiteralTemplate()
	fileTemplate()
}

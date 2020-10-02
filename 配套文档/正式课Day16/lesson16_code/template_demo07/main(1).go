package main

import (
	"log"
	"os"
	"html/template"
)

type User struct {
	FirstName 	string
	LastName	string
}

// 给结构体定义一个方法
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	t, err := template.ParseFiles("test.tpl")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	err = t.Execute(os.Stdout, User{FirstName: "ls", LastName: "lianshi"})
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}
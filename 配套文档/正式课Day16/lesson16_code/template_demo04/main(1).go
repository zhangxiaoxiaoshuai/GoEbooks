package main

import (
	"log"
	"os"
	"html/template"
)

type Item struct {
	Name	string
	Price	int
}

func main() {
	t, err := template.ParseFiles("test.tpl")
	if err != nil {
		log.Fatal("Parse error:", err)
	}

	items := []Item {
		{ "iPhone", 699 },
		{ "iPad", 799 },
		{ "iWatch", 199 },
		{ "MacBook", 999 },
	}

	err = t.Execute(os.Stdout, items)
	if err != nil {
		log.Fatal("Execute error:", err)
	}
}

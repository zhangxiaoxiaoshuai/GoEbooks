package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("index.tpl")
	if err != nil {
		fmt.Fprintf(w, "Parse file template error:%v", err)
		return
	}

	u := User{Name: "xiaop", Age: 18}
	err = t.Execute(w, u)
	if err != nil {
		fmt.Fprintf(w, "Parse file template error:%v", err)
	}
}


func main() {
	http.HandleFunc("/index", indexHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request){
	// 当用户以GET 请求 /login 时，表示用户想要一个登录的页面
	b, _ := ioutil.ReadFile("./login.html")
	w.Write(b)
}

func body(w http.ResponseWriter, r *http.Request){
	data := make([]byte, r.ContentLength)
	r.Body.Read(data) // 忽略错误处理
	defer r.Body.Close()

	fmt.Fprintln(w, string(data))
}


func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/body", body)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

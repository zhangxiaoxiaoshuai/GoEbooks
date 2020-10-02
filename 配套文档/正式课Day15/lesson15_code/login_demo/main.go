package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func login(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost{
		// 2. 当用户以 POST 方法 请求 /login 时，表示用户写把他的用户名和密码发送过来了，应该回复一个密码正确与否的响应
		// 2.1 从请求中获取用户填写的数据
		// r.Body  --> 读出所有的请求数据 --> JSON.Unmarshal  --> 结构体
		username := r.FormValue("username")
		email := r.FormValue("email")
		fileObj, fileHeader, err := r.FormFile("avatar")
		if err != nil {
			w.Write([]byte("upload file error"))
			return
		}
		b , _ := ioutil.ReadAll(fileObj)
		fmt.Println(fileHeader.Filename)
		ioutil.WriteFile(fileHeader.Filename, b, 0644)
		// 2.2 对数据进行校验
		fmt.Println(username, email)
		// 2.3 返回响应
		w.Write([]byte("ok"))
	}else{
		// 1. 当用户以 GET 方法 请求 /login 时，表示用户想要一个登录的页面，应该回一个页面的响应
		b, _ := ioutil.ReadFile("./login.html")
		w.Write(b)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request){
	rawQuery := r.URL.RawQuery
	fmt.Println(rawQuery)
	fmt.Fprintf(w, rawQuery)
}

func main() {
	// 当用户访问 /login 的时候，执行login函数
	http.HandleFunc("/login", login)
	http.HandleFunc("/search", searchHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

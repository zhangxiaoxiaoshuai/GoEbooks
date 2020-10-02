package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func loginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		username := r.FormValue("username")
		password := r.FormValue("password")
		// 2.2 对数据进行校验
		fmt.Println(username, password)
		if username == "xiaop" && password == "123456"{
			// 登陆成功
			// 创建Cookie数据
			c1 := &http.Cookie {
				Name:       "login",
				Value:      "true",
				HttpOnly:   true,
				// 默认关闭浏览器Cookie就失效
				MaxAge: 10,
			}
			// 把Cookie添加到响应头中
			//w.Header().Set("Set-Cookie", c1.String())
			http.SetCookie(w, c1)
			w.Write([]byte("登陆成功"))
		}else {
			// 登陆失败
			w.Write([]byte("登陆失败"))
		}
	}else{
		b, _ := ioutil.ReadFile("./login.html")
		w.Write(b)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	// 可以通过请求头中是否包含指定的 Cookie数据 来判断用户是否是登陆过的用户
	// 1. 如果是登录用户就返回个人中心页面

	c, err := r.Cookie("login")
	if err != nil {
		// 从当前请求中 取不到 login 这个cookie数据
		fmt.Println(err)
		// 2. 如果不是登录用户就跳转到登陆页面
		//  返回一个跳转到登陆页面的响应
		w.Header().Set("Location", "http://127.0.0.1:8080/login")
		w.WriteHeader(http.StatusFound)  // 必须作为修改响应头的最后一步
		return
	}
	fmt.Println(c.Value)
	fmt.Fprintf(w, "个人中心页面")
}

func main() {
	// 当用户访问 /login 的时候，执行login函数
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

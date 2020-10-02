package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Web!")
}

func taobao(w http.ResponseWriter, r *http.Request){
	// w:跟响应相关的
	// r:跟请求相关的
	fmt.Fprintf(w, "淘宝！")  // 写入响应的内容
}

func hi(w http.ResponseWriter, r *http.Request){
	// 1. 返回一个HTML文件
	// 1.1 打开文件，读取内容
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		b = []byte{}  // []int{}  []string{}
	}
	// 2. 返回数据
	w.Write(b)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/taobao", taobao)
	http.HandleFunc("/hi", hi)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method:%v\n", r.Method)
	URL := r.URL
	fmt.Fprintf(w, "Scheme: %s\n", URL.Scheme)
	fmt.Fprintf(w, "Host: %s\n", URL.Host)
	fmt.Fprintf(w, "Path: %s\n", URL.Path)
	fmt.Fprintf(w, "RawPath: %s\n", URL.RawPath)
	fmt.Fprintf(w, "RawQuery: %s\n", URL.RawQuery)
	fmt.Fprintf(w, "Fragment: %s\n", URL.Fragment)
	// Proto
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "ProtoMajor: %d\n", r.ProtoMajor)
	fmt.Fprintf(w, "ProtoMinor: %d\n", r.ProtoMinor)
	// Header
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s: %v\n", key, value)
	}
	// Body
	//  r.ContentLength -> 请求数据的长度
	data := make([]byte, r.ContentLength)
	r.Body.Read(data) // 从请求体中读取数据，忽略错误处理
	defer r.Body.Close()  // 一定要记得关闭r.Body

	fmt.Fprintln(w, string(data))
	fmt.Fprintf(w, "Hello, index!")
}


func main() {
	http.HandleFunc("/index", index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

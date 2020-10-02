package main

import (
	"fmt"
	"log"
	"net/http"
)

func writeHandler(w http.ResponseWriter, r *http.Request) {
	str := `<html>
		<head><title>Go Web</title></head>
		<body><h1>直接使用 Write 方法<h1></body>
		</html>`
	w.Write([]byte(str))
}

func writeHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "This API not implemented!!!")
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header() --> Header 类型 --> map[string][]string
	w.Header().Set("Location", "http://baidu.com")
	w.WriteHeader(http.StatusFound)  // 必须作为修改响应头的最后一步
}

func main() {

	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/writeheader", writeHeaderHandler)
	http.HandleFunc("/redirect", headerHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
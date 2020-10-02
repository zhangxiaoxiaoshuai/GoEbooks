package main

import (
	"fmt"
	"runtime"
)

//func init() {
//	var pcs [1]uintptr
//	//显示包下面具体函数名称
//	runtime.Callers(1, pcs[:])
//	fn := runtime.FuncForPC(pcs[0])
//	fmt.Println(fn.Name())
//}
//
//func init() {
//	var pcs [1]uintptr
//	runtime.Callers(1, pcs[:])
//	fn := runtime.FuncForPC(pcs[0])
//	fmt.Println(fn.Name())
//}
//  默认的init函数后缀
func init() {
	var pcs = make([]uintptr, 1)
	runtime.Callers(0, pcs[:])
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

func init() {
	var pcs = make([]uintptr, 1)
	runtime.Callers(0, pcs[:])
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

func main() {
	//_, file, _, _ := runtime.Caller(0)
	//fmt.Println(file)
}

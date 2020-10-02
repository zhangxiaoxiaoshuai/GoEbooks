package main

import (
	"fmt"
	"unsafe"
)

//type NULL struct {
//}

func main0801() {
	//如果定义多个空结构体都指向同一个内存地址 (全局区)
	//null := NULL{}
	null := struct{}{}
	fmt.Println(unsafe.Sizeof(null))
	fmt.Printf("%p\n", &null)

	null1:=struct{}{}
	fmt.Printf("%p\n", &null1)

}
type S struct {
	A struct{}
	B struct{}
}

func main(){
	s:=S{}
	fmt.Println(s.A)
	fmt.Println(s.B)
}
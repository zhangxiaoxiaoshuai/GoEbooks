package main

import (
	"fmt"
	"reflect"
)
//在结构体操作中 结合json时 结构体成员名首字母大写
type Perple struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//通过反射获取结构体成员
	typeof := reflect.TypeOf(Perple{})
	//遍历结构体成员
	for i := 0; i < typeof.NumField(); i++ {
		field := typeof.Field(i)
		tag := field.Tag.Get("json")
		fmt.Println(tag)
	}
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//空接口可以存储任意类型数据  空接口类型数据无法计算
	var i interface{} = 123
	var a interface{} = 234
	//fmt.Println(i + a)
	t1 := reflect.TypeOf(i)
	t2 := reflect.TypeOf(a).

	fmt.Printf("%T\n", t1)
	if t1 == t2 {
		v1 := reflect.ValueOf(i).Int()
		v2 := reflect.ValueOf(a).Int()
		fmt.Println(v1 + v2)
	}

}

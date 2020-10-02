package main

import (
	"fmt"
	"sync"
)

func Print(key, value interface{}) bool {
	fmt.Println("键：", key, "值：", value)
	return true
}
func main() {
	//var m map[interface{}]interface{}//err map必须初始化才可以使用 make
	//m["法师"] = "帅哥"
	//fmt.Println(m)
	var smap sync.Map
	//将键值对存储在sync.map
	smap.Store("法师", 33)
	smap.Store("兵哥", 18)
	smap.Store("木子", 22)

	//从map中获取数据
	//value, ok := smap.Load("老子")
	//if ok {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("未找到数据")
	//}

	//从map中删除数据
	smap.Delete("法师")

	//遍历数据(函数回调  将函数作为函数参数)
	smap.Range(Print)
}

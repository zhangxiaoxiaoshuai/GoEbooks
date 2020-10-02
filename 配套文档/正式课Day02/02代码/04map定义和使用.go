package main

import (
	"fmt"
	"unsafe"
)

func main0401() {
	//map定义
	//var 字典名 map[键类型]值类型
	//var m map[int]string = make(map[int]string, 10)
	//m[1001] = "法师"
	//m[8888] = "兵哥"
	//
	//fmt.Println(m)

	//map是一个无序的的集合
	m := map[int]string{1001: "法师", 8888: "兵哥", 1234: "兴臣", 3333: "孟跃平"}
	//fmt.Println(m)

	//k key 键 v value 值
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func main() {
	//map中的key必须支持== ！=运算符操作  不能是切片 map 函数
	m := make(map[int]string)
	m[1001] = "法师"
	m[1005] = "兵哥"

	//通过key判断value是否存在
	if v, ok := m[1001]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在key")
	}

	fmt.Println(len(m))
	//删除map中的数据
	delete(m, 1001)
	fmt.Println(len(m))

	fmt.Println(m)

	fmt.Println(unsafe.Sizeof(m))

}

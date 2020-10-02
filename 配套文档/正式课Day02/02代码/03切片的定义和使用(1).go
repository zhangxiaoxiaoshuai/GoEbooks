package main

import (
	"fmt"
	"unsafe"
)

func main0301() {
	//切片的定义和使用
	//var 切片名 []数据类型
	//var slice []int
	//make([]数据类型，长度)
	//var slice []int = make([]int, 10)
	//slice[0] = 123
	//fmt.Println(slice)

	var slice []int
	fmt.Println(len(slice)) //计算长度
	fmt.Println(cap(slice)) //计算容量

	//使用append对切片进行扩容
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(len(slice)) //计算长度
	fmt.Println(cap(slice)) //计算容量

	//切片扩容为偶数的值  小于1024 为上一次的两倍 大于1024为上一次的1/4
	fmt.Println(slice)
}

func main0302() {
	//切片的截取
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//切片名[起始下标：结束下标：容量]  左闭右开 包含起始下标 不包含结束下标
	//s := slice[2:7]
	//fmt.Println(s)
	//s:=slice[2:]
	//fmt.Println(s)
	//s:=slice[:5]
	//fmt.Println(s)
	s := slice[2:5:6] //实际容量=容量-起始下标
	fmt.Println(s)

	//fmt.Println(len(s))
	//fmt.Println(cap(s))

	s[0] = 333
	//切片的截取 是将新的切片指向源切片的内存地址  修改一个会影响另外一个
	fmt.Println(s)
	fmt.Println(slice)

	fmt.Println(unsafe.Sizeof(slice))
	fmt.Println(unsafe.Sizeof(s))
}
func main() {
	slice := []int{1, 2, 3, 4, 5}

	//s := slice
	//s[2] = 333
	//fmt.Println(s)
	//fmt.Println(slice)
	//在存储存储两个内容完全相同 但是不会相互影响
	s := make([]int, 5)
	copy(s, slice)
	s[2]=333
	fmt.Println(s)
	fmt.Println(slice)

}

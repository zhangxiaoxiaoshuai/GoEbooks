package main

import (
	"fmt"
	"unsafe"
)

func Test(slice []int) []int {
	slice = append(slice, 1, 2, 3, 4, 5)

	//fmt.Println("demo函数：", slice)
	//slice[0] = 666
	//slice[1] = 777
	//slice[2] = 888
	//slice[3] = 999
	//slice[4] = 000
	return slice
}
func main0301() {
	slice := []int{1, 2, 3, 4, 5}

	slice = Test(slice)

	fmt.Println(slice)
}

func Demo(slice *[]int) {
	*slice = append(*slice, 1, 2, 3, 4, 5)
}
func main() {
	/*
	//所有内存地址 是一个无符号十六进制整型数据表示的
	type slice struct {
		array unsafe.Pointer	//指针 8字节
		len   int				//长度 8字节
		cap   int				//容量 8字节
	}
	 */
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(slice))
	Demo(&slice)
	fmt.Println(unsafe.Sizeof(slice))

	fmt.Println(slice)
}

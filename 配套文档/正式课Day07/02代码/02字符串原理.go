package main

import "fmt"

func main0201() {
	str := "hello world"
	//fmt.Println(unsafe.Sizeof(str))
	//字符串下标读取字符
	//fmt.Println(str[2])
	//字符串下标修改字符
	//str[2]='M'//err

	//创建切片
	slice := []byte(str)
	//修改切片中的元素的值
	slice[2] = 'M'
	fmt.Println(slice)
}

func main() {
	slice := []byte{'h', 'e', 'l', 'l', 'o'}
	str := string(slice) //如果出现赋值会发生内存拷贝

	//str[2]='M'//err 字符串常量区不允许修改

	//修改slice 不会影响str
	slice[2] = 'M'
	fmt.Println(str)
	fmt.Println(string(slice)) //不会内存拷贝

	//硬编码常量

	//fmt.Println("helloworld")
	str1 := "helloworld"
	fmt.Println(str1 == "helloworl")
}

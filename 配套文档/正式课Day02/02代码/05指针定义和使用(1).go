package main

import "fmt"

func main0501() {
	a := 10
	//取出变量a在内存的地址
	//fmt.Println(&a)
	var p *int = &a
	//fmt.Println(p)
	//fmt.Println(&a)
	//通过指针间接修改变量的值
	*p = 132
	fmt.Println(a)

	//const MAX int = 100
	//fmt.Println(&MAX)//err 不可以获取常量的内存地址
}
func main() {
	//定义指针 默认值为nil 指向内存地址编号为0的空间  内存地址0-255为系统占用 不允许用户读写操作
	//var p *int = nil
	//*p = 123 //err
	//fmt.Println(*p)

	//开辟数据类型大小的空间 返回值为指针类型
	//new(数据类型)
	var p *int
	p = new(int)
	*p=123
	fmt.Println(*p)

}

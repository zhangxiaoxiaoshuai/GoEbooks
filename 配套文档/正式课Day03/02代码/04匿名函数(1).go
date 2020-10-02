package main

import "fmt"

//func test(){
//
//}
//定义函数类型
type FUNCTYPE func(int, int)

func main() {
	//匿名函数
	//func(a int, b int) {
	//	fmt.Println(a + b)
	//}(10,20)

	//定义函数类型变量f
	//var f FUNCTYPE
	//var f func(int,int)
	f := func(a int, b int) {
		fmt.Println(a + b)
	}

	fmt.Printf("%T\n",f)
	//函数在代码区的内存地址
	//fmt.Println(test)
	f(10, 20)
	//fmt.Println(f)
}

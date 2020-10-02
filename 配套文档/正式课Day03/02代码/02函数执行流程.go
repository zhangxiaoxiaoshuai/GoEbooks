package main

import "fmt"

//函数定义
func add(a int, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 20
	//函数调用
	value := add(a, b)
	fmt.Println(value)
}

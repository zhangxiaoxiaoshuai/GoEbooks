package main

import (
	"errors"
	"fmt"
)

/*
编辑时异常
编译时异常
运行时异常
*/
func demo(i int) {
	//recover错误拦截  出现在panic错误之前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var arr [10]int
	//数组下标越界
	arr[i] = 123 //err
	fmt.Println(arr)
}
func main0501() {

	demo(10)
	fmt.Println("程序继续执行...")
}

//一般性错误处理
func test(a int, b int) (v int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	v = a / b
	return
}
func main() {
	test(10, 0)
}

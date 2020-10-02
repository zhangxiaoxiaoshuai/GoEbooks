package main

import "fmt"

func main() {
	//defer 延迟调用 将函数加载到内存中并没有执行 而是在函数出栈时在执行
	defer fmt.Println("性感法师")
	defer fmt.Println("在线教学")
	defer fmt.Println("轻松就业")
	defer fmt.Println("日薪越亿")
}

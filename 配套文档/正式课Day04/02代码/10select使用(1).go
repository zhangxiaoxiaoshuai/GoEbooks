package main

import (
	"fmt"
	"runtime"
	"time"
)

func main1001() {
	//指定CPU核数
	//fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(1)
	intch := make(chan int, 1)
	strch := make(chan string, 1)
	intch <- 1
	strch <- "hello"
	/*
		1、select case如果未准备好 会阻塞
		2、如果多个case都满足条件会随机选择一个执行
		3、case 必须为IO操作
		4、可以做超时和退出判断
		5、select  如果对应的是异步处理 需要在for循环中使用
		6、select 可以监控chan数据流向
	*/
	select {
	case value := <-intch:
		fmt.Println(value)
	case value := <-strch:
		fmt.Println(value)
	case <-time.After(time.Second * 5): //超时处理
		fmt.Println("超时")
		//default:阻塞操作
	}
}
func main() {
	timer := time.After(time.Second)
	fmt.Printf("%T\n", timer)
}

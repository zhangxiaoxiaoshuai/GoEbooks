package main

import (
	"fmt"
	"time"
)

func main() {
	//chan定义方式
	// 变量名 chan=make(chan 数据类型,大小)
	//ch := make(chan int, 10)//双向带缓冲区
	//fmt.Println(unsafe.Sizeof(ch))

	//单向chan
	//ch:=make(chan<- int) //接收
	//ch := make(<-chan int) //发送

	//有缓冲和无缓冲区别
	//无缓冲 要求发送和接收必须同时准备好 才可以完成  属于同步操作 会发生阻塞等待
	//有缓冲 可以接收N个值 属于异步操作
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	ch <- 1
	time.Sleep(time.Second * 2)
}

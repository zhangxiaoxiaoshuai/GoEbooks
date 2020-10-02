package main

import (
	"fmt"
	"sync"
	"time"
)

func Test() {
	fmt.Println("性感法师在线教学~")
}
func main() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		//多次在协程调用  执行一次
		go func() {
			once.Do(Test) //函数回调
		}()
		//Test()
	}

	time.Sleep(time.Second * 2)
}

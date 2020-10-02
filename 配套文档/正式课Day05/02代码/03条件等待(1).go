package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	cond := sync.NewCond(new(sync.Mutex))

	for i := 0; i < 3; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			fmt.Println("协程启动：", i)
			cond.L.Lock() //加锁
			fmt.Println("协程：", i, "加锁")
			cond.Wait()

			cond.L.Unlock() //解锁
			fmt.Println("协程：", i, "解锁")

		}(i)
	}

	time.Sleep(time.Second * 2)
	//发送信号
	cond.Signal()
	fmt.Println("主协程发送信号")

	time.Sleep(time.Second * 2)
	//发送信号
	cond.Signal()
	fmt.Println("主协程发送信号")

	time.Sleep(time.Second * 2)
	//发送信号
	cond.Signal()
	fmt.Println("主协程发送信号")

	wg.Wait()
}

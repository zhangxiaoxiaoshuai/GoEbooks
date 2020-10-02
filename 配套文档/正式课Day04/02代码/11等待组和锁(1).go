package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	wg := sync.WaitGroup{}
	count := 0
	//add 添加 done 完成（add(-1)）wait
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			m.Lock()         //加锁

			defer wg.Done()  //完成
			//数据处理
			count++


			m.Unlock() //解锁
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println("程序继续执行")

}

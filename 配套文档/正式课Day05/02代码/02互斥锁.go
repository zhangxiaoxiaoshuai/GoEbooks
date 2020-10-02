package main

import (
	"fmt"
	"sync"
)

//不建议全局变量在协程中使用  如果使用 需要加锁
var num = 0
var wg sync.WaitGroup
var mu sync.Mutex

func Count() {
	mu.Lock()
	defer mu.Unlock()
	num++
	wg.Done()
}
func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go Count()
	}

	wg.Wait()

	fmt.Println(num)
	//sync.RWMutex{}
}

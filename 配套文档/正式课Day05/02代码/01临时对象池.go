package main

import (
	"fmt"
	"sync"
)
//定义函数赋值给pool.New
func Demo() interface{} {
	return 3.14
}

func main() {
	var pool sync.Pool
	//put 将数据存储在临时对象池
	pool.Put(1)
	pool.Put("hello")
	pool.Put(true)
	pool.Put(3.14)

	//get 将数据从临时对象池取出
	value := pool.Get()
	fmt.Println(value)
	//临时对象池第一个数据在最前，后续的数据采用先进后出的原则
	value = pool.Get()
	fmt.Println(value)

	//循环获取所有对象池内容
	for {
		value := pool.Get()
		if value == nil {
			break
		}
		fmt.Println(value)
	}
	//需要函数类型变量
	pool.New = Demo
}

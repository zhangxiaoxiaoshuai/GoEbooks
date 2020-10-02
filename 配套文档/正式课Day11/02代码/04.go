package main

import "fmt"

func main() {

	list := make([]int, 0)
	//list := new([]int) //list数据类型：*[]int
	//fmt.Printf("%T\n",list)
	//list = append(list, 1)

	//对指针类型进行解引用
	//*list = append(*list, 1, 2, 3)
	fmt.Println(list)
}

package main

import "fmt"

func main()  {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()

	panic("触发异常")
}
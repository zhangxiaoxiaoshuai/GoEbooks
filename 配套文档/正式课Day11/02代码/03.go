package main

import "fmt"

type Test struct {
	Name string
}

var list map[string]Test

func main() {

	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	//list["name"].Name//可读 不可以
	//list["张三"].Name = "Hello"
	fmt.Println(list["name"].Name)
}

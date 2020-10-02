package main

import "fmt"

type people struct {
	Name string
}

func main0501() {
	var i interface{} = &people{"木子"}
	//类型断言
	if v, ok := i.(*people); ok {
		fmt.Println(v.Name)
	} else {
		fmt.Println("类型匹配失败")
	}
}
func main() {
	var i interface{} = people{"法师"}
	switch i.(type) {
	case people:
		fmt.Println("perple")
	case *people:
		fmt.Println("*people")
	default:
		fmt.Println("类型匹配失败")

	}
}

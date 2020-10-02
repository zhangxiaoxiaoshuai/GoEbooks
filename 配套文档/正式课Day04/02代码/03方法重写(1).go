package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}
type Student1 struct {
	Person1
	Score int
}

func (p *Person1) Print() {
	fmt.Println(*p)
}
func (s *Student1) Print() {
	fmt.Println(*s)
}

func main() {
	stu := Student1{Person1{"木子", 18}, 100}
	//采用就近原则 默认使用本结构体对应的方法
	//stu.Print()
	//使用父类的方法
	stu.Person1.Print()
}

package main

import "fmt"

//子类和父类有共同的成员名
type Perple1 struct {
	Name string
	Age  int
}

type Student1 struct {
	//P Perple1
	//匿名字段 将结构体作为另外一个结构体成员
	Perple1
	Name  string
	Score int
}

func main() {
	stu := Student1{Perple1{"兵哥", 18}, "兵哥哥", 100}
	//stu.Name = "法师"
	//stu.P.Name = "法师"

	//采用就近原则 修改本结构体信息
	//stu.Name="法师"
	stu.Perple1.Name="法师"
	fmt.Println(stu)
}

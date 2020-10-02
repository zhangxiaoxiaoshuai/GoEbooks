package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type student struct {
	Person
	Score int
}

func (p *Person) SayHello() {
	fmt.Println(*p)

}

func main() {
	//创建stu对象
	stu := student{Person{"兵哥", 18}, 100}
	//子类可以直接使用父类的方法
	stu.SayHello()
}

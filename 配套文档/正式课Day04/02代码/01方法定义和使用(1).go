package main

import "fmt"

//定义结构体表示对象属性
type Student struct {
	Name string
	age  int
}

//对象行为用方法表示
//func (方法接收者)方法名（参数）（返回值）{return 值}
//值对象接收者
//func (stu Student) SayHello() {
//	//不可以修改对象的属性
//	stu.Name = "法师"
//	//fmt.Println("大叫好", "我是", stu.Name)
//	fmt.Println(stu)
//}

//指针对象接收者
func (stu *Student) SayHello() {
	//fmt.Println("大叫好", "我是", stu.Name)
	//可以修改对象属性
	stu.Name = "法师"
}
func main() {
	stu := Student{"兵哥", 18}
	stu.SayHello()
	//(&stu).SayHello()
	fmt.Println(stu)
	fmt.Println(stu.SayHello)
}

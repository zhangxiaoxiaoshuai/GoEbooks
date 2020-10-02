package main

import "fmt"

/*
type 接口名 interface{
	方法声明
	方法名（参数）返回值
}
*/

type Humaner interface {
	SayHi()
}

type Person2 struct {
	Name string
	Age  int
}

type Teacher struct {
	Person2
	Subject string
}
type Student2 struct {
	Person2
	Score int
}

func (t *Teacher) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我教%s\n", t.Name, t.Age, t.Subject)
}
func (s *Student2) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我的成绩%d\n", s.Name, s.Age, s.Score)
}

func main0401() {
	//tea := Teacher{Person2{"法师", 33}, "go语言开发"}
	stu := Student2{Person2{"木子", 18}, 100}
	//创建接口类型变量
	var h Humaner
	//将对象的地址赋值给接口类型
	h = &stu
	//通过接口调用
	h.SayHi()
}

//多态 将接口类型作为函数参数
func SayHi(h Humaner) {
	h.SayHi()
}
func main() {
	//初始化对象
	//tea := Teacher{Person2{"法师", 33}, "go语言开发"}
	stu := Student2{Person2{"木子", 18}, 100}
	//test:= struct {}{}
	SayHi(&stu)
}

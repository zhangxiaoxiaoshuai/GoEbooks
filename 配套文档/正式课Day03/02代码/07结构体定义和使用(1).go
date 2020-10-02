package main

import "fmt"

type student struct {
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main0701() {
	//定义结构体变量
	stu := student{"兵哥哥", 18, "男", 100, "山东聊城"}
	//stu.sex = "女"
	//fmt.Println(stu)
	fmt.Printf("%p\n", &stu)
	fmt.Println(&stu.name)
	fmt.Println(&stu.age)
	fmt.Println(&stu.sex)
	fmt.Println(&stu.score)
	fmt.Println(&stu.addr)
}

type Student struct {
	name string
	age  int
}

func main() {
	m := make(map[int]Student)
	s1 := Student{"兵哥", 18}
	s2 := Student{"法师", 33}
	m[1001] = s1

	//m[1001].name = "法师" //写操作 err
	m[1001] = s2 //如果想修改结构体  必须整体操作
	fmt.Println(m)
	fmt.Println(m[1001].name) //读操作 OK
}

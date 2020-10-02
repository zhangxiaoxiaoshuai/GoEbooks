package main

/*
int Add(int a, int b)
{
	return a + b;
}
*/
import "C"
import "fmt"

func main() {
	a := 10
	b := 20
	//将go语言数据类型转成C语言数据类型
	value := int(C.Add(C.int(a), C.int(b)))
	fmt.Println(value)
	fmt.Printf("%T\n", value)
	//强制类型转换
	//c := int(value)
	//fmt.Println(c)
}

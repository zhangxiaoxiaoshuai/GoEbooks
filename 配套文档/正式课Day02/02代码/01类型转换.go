package main

import "fmt"

func main0101() {
	//去市场买菜 黄瓜3.25元/斤 买5斤
	price := 3.25 //float64
	weight := 5   //int

	//fmt.Printf("%T\n", price)
	//fmt.Printf("%T\n", weight)
	//数据类型（变量名）
	//sum := price * float64(weight)//16.225
	//将浮点型转成整型 可能丢失数据精度
	//sum := int(price) * weight //15

	//数据类型（表达式）
	sum := int(price * float64(weight)) //16
	fmt.Println(sum)
}

type int1 int
type int2 int

func main() {
	//var a int8 = 123
	//var b int = 234
	//fmt.Println(int(a) + b)
	//type rune = int32  两个数据类型可以计算
	//ch := 'A' //rune ==int32
	//var b int32 = 100
	//fmt.Println(ch + b)
	var a int1 = 123
	var b int2 = 123

	fmt.Println(a + int1(b))
}

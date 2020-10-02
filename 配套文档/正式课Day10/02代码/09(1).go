package main

import "fmt"

func main0901() {
	a := 2 ^ 15
	b := 4 ^ 15

	/*
		0010
		1111
		1101=13

		0100
		1111
		1011=11
	*/

	if a > b {
		println("a:", a)
	} else {
		println("b:", b)
	}
}

func main() {
	a := 10
	b := -20

	//a,b=b,a
	//交换变量的值
	a = a ^ b
	b = a ^ b
	a = a ^ b


	fmt.Println(a)
	fmt.Println(b)
}

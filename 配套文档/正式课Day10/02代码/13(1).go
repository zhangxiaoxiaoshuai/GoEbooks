package main

import "fmt"

func main() {
	a := []int{0,1,2,3,4,5,6,7}
	b := a[:3]
	b = append(b,4)
	fmt.Println(a)
	fmt.Println(b)
}
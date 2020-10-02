package main

import (
	"fmt"
)

func f(a int, b uint) {
	val := copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("%d\n", val)
}

func main() {
	f(90, 314)
	//slice := []string{}
	//fmt.Println(unsafe.Sizeof(slice))
}

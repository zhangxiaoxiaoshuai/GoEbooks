package main

import (
	"fmt"
	"unsafe"
)

type Semple struct {
	//会按照操作系统的位数进行对齐
	a int8   //1
	c uint16 //2
	b int    //8
}

func main() {
	s := Semple{1, 2, 3}

	fmt.Println(unsafe.Sizeof(s))
}

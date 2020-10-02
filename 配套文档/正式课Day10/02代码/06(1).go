package main

import "fmt"

func main() {
	//a1 := 100.5        //float64
	//a2 := 2147483447.0 //float64
	//var x uint32 = a1 + a1 + a2

	var x uint32 = 100.5 + 100.5 + 2147483447.0
	//unsigned int32 0~ 4294967295（2^32）-1
	fmt.Println(x)
}

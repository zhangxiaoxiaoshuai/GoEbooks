package main

import "runtime"

func main() {
	//fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(1)
}

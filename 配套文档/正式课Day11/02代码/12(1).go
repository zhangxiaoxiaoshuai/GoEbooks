package main

type T []int

func F(t T) {}

func main() {
	var x []int
	var y T
	y = x
	F(x)
	_ = y
}

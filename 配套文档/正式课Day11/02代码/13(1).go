package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		a := i
		funs = append(funs, func() {
			println(&a, a)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

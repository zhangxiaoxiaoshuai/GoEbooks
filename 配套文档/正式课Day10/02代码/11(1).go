package main

import "fmt"

type Hash [32]byte

func MustNotBeZero(h Hash) {
	if h == [32]byte{} {
		fmt.Println(1)
	}
}

func main() {
	MustNotBeZero(Hash{})
}
package main

import (
	"fmt"
)

func main() {
	sn1 := struct {
		age  int
		m   map[string]string
	}{age: 11,m: map[string]string{"a": "1"}}
	sn2 := struct {
		age  int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//if sn1 == sn2 {
	//	fmt.Println("sn1 == sn2")
	//}
}
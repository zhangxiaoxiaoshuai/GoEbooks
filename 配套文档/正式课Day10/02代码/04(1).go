package main

import "fmt"

func A(string string) string {
	return string + string
}
func B(len int) int {
	return len+len
}
//func C(val, default string) string {
//	if val == "" {
//		return default
//}
//return val
//}
func main() {
	val:=A("hello")
	fmt.Println(val)
}

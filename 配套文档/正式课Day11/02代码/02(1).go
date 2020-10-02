package main

import "fmt"

//func main() {
//	m := make(map[string]int)
//	m["foo"]++	//m["foo"]=m["foo"] + 1
//	fmt.Println(m["foo"])
//}

func main() {
	//统计字符串中字符出现的个数
	str := "sdjnfmnjghfhfhjdgdjhsfgjhjhddhjdgjfybjdjghjcjhdhfhgjgh"

	m := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}

	//fmt.Println(m)

	for k, v := range m {
		fmt.Printf("%c:%d\n", k, v)
	}

}

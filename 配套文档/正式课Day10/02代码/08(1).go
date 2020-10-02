package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.TrimRight("abcdefedcmba", "fedabc")
	fmt.Printf("%q\n", s)
}

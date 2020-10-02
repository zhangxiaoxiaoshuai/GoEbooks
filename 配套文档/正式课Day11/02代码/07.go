package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := GetValue()

	fmt.Println(reflect.TypeOf(i)) //foat64

	//只有接口类型可以使用类型断言
	//如果在case中出现真实类型和interface{}谁在前谁有限级别高
	switch i.(type) {
	case float64:
		println("float64")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")

	}
}
func GetValue() interface{} {
	return 1.23
}


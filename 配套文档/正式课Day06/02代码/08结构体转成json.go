package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score []int
	Addr  string
}

func main() {
	stu := Student{"法师", 33, []int{100, 100, 59}, "黑龙江大庆"}
	//生成json数据格式
	//slice, err := json.Marshal(stu)
	//生成格式化json数据格式
	slice, err := json.MarshalIndent(stu, "", "   ")
	if err != nil {
		fmt.Println("json格式转换失败")
	} else {
		//fmt.Println(string(slice))
		fmt.Println(slice)
	}

}

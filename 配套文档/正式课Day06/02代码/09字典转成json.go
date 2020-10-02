package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建字典保存键值对信息
	m := make(map[string]interface{})
	//数据存储
	m["Name"] = "法师"
	m["Age"] = 33
	m["Score"] = []int{100, 100, 59}
	m["Addr"] = "黑龙江大庆"

	//slice, err := json.Marshal(m)

	//格式化转成json
	slice, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		fmt.Println("json格式转换失败")
	} else {
		fmt.Println(string(slice))
	}
}

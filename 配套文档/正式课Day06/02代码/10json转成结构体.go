package main

import (
	"encoding/json"
	"fmt"
)

type Student1 struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Score []int  `json:"Score"`
	Addr  string `json:"Addr"`
}

func main() {
	slice := []byte(`
{
    "Addr": "黑龙江大庆",
    "Age": 33,
    "Name": "法师",
    "Score": [
       100,
       100,
       59
    ]
 }
`)
	var stu Student1
	err := json.Unmarshal(slice, &stu)

	if err != nil {
		fmt.Println("json格式转换失败")
	} else {
		fmt.Println(stu)
	}
}

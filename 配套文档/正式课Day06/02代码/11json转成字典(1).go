package main

import (
	"encoding/json"
	"fmt"
)

func main() {
//	slice := []byte(`
//{
//    "Addr": "黑龙江大庆",
//    "Age": 33,
//    "Name": "法师",
//    "Score": [
//       100,
//       100,
//       59
//    ]
// }
//`)
	//slice:=[]byte("{\"Addr\":\"黑龙江大庆\",\"Age\":33,\"Name\":\"法师\",\"Score\":[100,100,59]}")
	slice:=[]byte(`{"Addr":"黑龙江大庆","Age":33.3,"Name":"法师","Score":[100,100,59]}`)

	var i interface{}
	err := json.Unmarshal(slice, &i)
	if err != nil {
		fmt.Println("json格式转换失败")
	} else {
		fmt.Println(i)
		//map[Addr:黑龙江大庆 Age:33 Name:法师 Score:[100 100 59]]
	}
	m := i.(map[string]interface{})
	//fmt.Println(m)
	//fmt.Printf("%T\n",m)
	for _, v := range m {
		//fmt.Println(k, v)
		//fmt.Printf("%T\n", v)
		switch val := v.(type) {

		case string:
			fmt.Println("string:", val)
		case float64:
			fmt.Println("float64", val)
		case []interface{}:
			//fmt.Println("[]int", val)
			for i, value := range val {
				fmt.Println(i, value)
				fmt.Printf("%T\n",value)
			}
		}
	}

}

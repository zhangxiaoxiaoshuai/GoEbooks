package main

import "fmt"

func main() {
	//if 表达式 {代码}  如果表达式为真 执行{}的代码
	//if 表达式 {代码1} else{代码2} 如果表达式为真 执行代码1内容 如果为假 执行代码2
	//if 表达式1 {代码1} else if表达式2{代码2}else{代码3} 依次判断表达式 满足条件执行

	score := 0
	//通过键盘获取数据  存放变量中
	fmt.Scan(&score) //阻塞式请求

	if score > 700 {
		fmt.Println("我要上清华")
	} else if score > 680 {
		fmt.Println("我要上北大")
	} else {
		fmt.Println("我要上蓝翔")
	}
}

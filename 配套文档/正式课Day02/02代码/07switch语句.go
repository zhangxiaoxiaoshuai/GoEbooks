package main

import "fmt"

func main0701() {
	//switc 表达式{case 值1：代码1 case 值2：代码2} 那个分支满足条件执行那一个

	//输入月份 计算天数
	day := 0
	m := 0
	fmt.Scan(&m)
	switch m {
	//如果有多个值 执行相同代码 可以使用逗号分隔
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
		//fallthrough //执行完当前分支 继续向下执行
	case 2:
		day = 29
	case 4, 6, 9, 11:
		day = 30
	default:
		//默认选择
		day = -1
	}
	fmt.Println(day)
}
func main0702() {
	//pi := 3.14
	////浮点型因为精度问题可能在switc中认为是同一个值
	//switch pi {
	//case 3.140000000000000000019999999:
	//	fmt.Println(pi)
	//case 3.14:
	//	fmt.Println(pi)
	//}
}
func main() {
	score := 0
	fmt.Scan(&score)
	switch {
	case score > 700:
		fmt.Println("我要上清华")
	case score > 680:
		fmt.Println("我要上北大")
	default:
		fmt.Println("分数错误，程序异常")
	}
}

package main

import "fmt"

func main0201() {
	//数组定义和使用
	//var 数组名 [元素个数]数据类型
	//var arr [10]int//默认值为0
	////使用数组名+下标进行数组初始化   下标是从0开始的到数组最大元素个数-1
	//arr[0] = 123
	//arr[-1] = 123 //err 数组下标越界

	//var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(arr)
	//len(数组名) 计算数组元素个数
	//fmt.Println(len(arr))

	//遍历数组元素
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	//i index 下标 v value 值
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		//& 取地址运算符
		fmt.Println(&arr[i])
	}
}

package main

import "fmt"

func main0801() {
	/*
		for 表达式1；表达式2；表达式3{
			代码体
		}
		for 表达式 {
			代码体
		}
		for range 集合{
			代码体
		}
	*/

	//计算1-100和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func main0802() {
	//嵌套循环
	//外层执行一次 内层执行一周
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i, j)
		}
	}
}

func main() {
	//冒泡排序
	//比较两个相邻数据元素 满足条件 交换数据位置

	slice := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}

	//外层控制行 表示执行次数
	for i := 0; i < len(slice)-1; i++ {
		//内层控制列 表示比较次数
		for j := 0; j < len(slice)-1-i; j++ {
			//比较两个相邻元素  大于号 升序排序  小于号 降序排序
			if slice[j] < slice[j+1] {
				//交换数据的值  = 多重赋值
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)

}

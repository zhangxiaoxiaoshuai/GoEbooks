package main

import "fmt"

func main0101() {
	//break 跳出循环
	//i:=0
	//for  {
	//	i++
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 5 {
				break
			}
			fmt.Println(i, j)
		}
	}
}
func main0102() {
	//continue 结束本次循环 继续下次循环
	//计算1-100偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		//奇数
		if i%2 == 1 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}

func main0103() {
	//goto 无条件跳转语句
	fmt.Println("性感法师")
FLAG:
	fmt.Println("在线讲课")
	fmt.Println("轻松就业")
	goto FLAG
	fmt.Println("日薪越亿")
}
func main() {
	//goto 可不以跳入循环内部  不可以跳出函数
HERE:
	for i := 0; i <= 10; i++ {
		if i == 5 {
			//break
			goto HERE
		}
		fmt.Println(i)
	}
	fmt.Println("程序执行结束...")
}

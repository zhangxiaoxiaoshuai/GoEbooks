package main

import (
	"fmt"
	"os"
)

//文件创建
func main1201() {
	fp, err := os.Create("D:/a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	//延迟关闭文件
	defer fp.Close()

	fmt.Println("文件创建成功")

}

//文件打开
func main() {
	//只读方式打开文件
	//os.Open()
	//打开文件同时设置权限
	fp, err := os.OpenFile("D:/a.txt", os.O_RDWR|os.O_CREATE, 600)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	//关闭文件
	defer fp.Close()
	//写入字符串
	fp.WriteString("法师真胖~")
}

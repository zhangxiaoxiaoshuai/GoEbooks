package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//程序运行时用户传递的参数信息
	args := os.Args
	fmt.Println(args)

	if args == nil || len(args) != 3 {
		fmt.Println("复制文件失败")
		return
	}
	//copy.exe 源文件 目标文件
	src := args[1]
	dst := args[2]

	if src == dst {
		fmt.Println("源文件和目标文件相同")
		return
	}

	fp1, err := os.Open(src)
	if err != nil {
		fmt.Println("源文件操作失败")
		return
	}
	fp2, err := os.Create(dst)
	if err != nil {
		fmt.Println("目标文件操作失败")
		return
	}
	defer fp1.Close()
	defer fp2.Close()

	//拷贝
	buf := make([]byte, 1024)//1kb
	for {
		//读取源文件内容
		n, err := fp1.Read(buf)
		if err != nil && err == io.EOF {
			break
		}
		//将有效数据写入在目标文件中
		fp2.Write(buf[:n])

	}

}

package main

import (
	"fmt"
	"io"
	"os"
)

func main1401() {
	fp, err := os.Open("D:/b.txt")
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer fp.Close()

	buf := make([]byte, 1023)
	//fp.Read(buf)
	//fmt.Println(string(buf))

	//读取整个文件
	for {
		n, err := fp.Read(buf)
		//io.EOF 文件结尾标志
		if err != nil && err == io.EOF {
			//fmt.Println(err)
			break
		}
		fmt.Print(string(buf[:n]))
	}

}
func main() {
	fp, err := os.Open("D:/b.txt")
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer fp.Close()
	buf := make([]byte, 6)
	fp.ReadAt(buf,12)
	fmt.Println(string(buf))
}

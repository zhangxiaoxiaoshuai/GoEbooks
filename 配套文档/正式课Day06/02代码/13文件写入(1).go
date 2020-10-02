package main

import (
	"fmt"
	"io"
	"os"
)

func main1301() {
	fp, err := os.OpenFile("D:/b.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("文件操作失败:", err)
		return
	}
	defer fp.Close()
	//在windows中换行是以\r\n为操作符 在linux中换行是以\n为操作符
	//slice := []byte("锄禾日当午\r\n汗滴禾下土")
	////写入文件
	//fp.Write(slice)

	str := "日照香炉生紫烟"
	fp.WriteString(str)
}
func main() {
	fp, err := os.OpenFile("D:/b.txt", os.O_WRONLY, 600)
	if err != nil {
		fmt.Println("文件操作失败：", err)
		return
	}
	defer fp.Close()
	//会覆盖偏移后的内容
	fp.WriteAt([]byte("李白"), 12)
	index, _ := fp.Seek(12, io.SeekStart)
	fmt.Println(index)
}

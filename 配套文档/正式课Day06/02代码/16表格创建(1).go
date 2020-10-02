package main

import (
	"fmt"
	"github/tealeg/xlsx"
)

func main() {
	//创建表格文件
	fp := xlsx.NewFile()
	sheet, err := fp.AddSheet("第一页")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i <= 10; i++ {
		row := sheet.AddRow()
		for j := 0; j <= 5; j++{
			cell := row.AddCell()
			cell.Value = "姓名"
		}
	}

	//保存表格文件
	err = fp.Save("D:/test.xlsx")
	if err != nil {
		fmt.Println("表格文件保存失败", err.Error())
	}
}

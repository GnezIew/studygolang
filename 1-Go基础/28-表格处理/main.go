package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	//新建表格
	file :=xlsx.NewFile()
	//保存文件
	err := file.Save("E:/Go/test.xlsx")
	if err != nil{
		fmt.Println("表格文件有误")
		return
	}
}

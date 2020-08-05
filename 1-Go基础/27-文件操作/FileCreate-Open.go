package main

import (
	"fmt"
	"os"
)

func CreateFile() {
	//创建文件
	f, err := os.Create("./demo.txt")
	if err != nil {
		fmt.Println("Create err :", err)
		return
	}

	//关闭文件
	defer f.Close()
	fmt.Println("create successful")
}

func Openfile() {
	//os.Open 只读方式打开文件
	f, err := os.Open("E:/Go/demo.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	defer f.Close()
	//操作文件
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]))
}

func Open_Writefile() {
	f, err := os.OpenFile("E:/Go/demo.txt", os.O_RDWR|os.O_CREATE, 666)
	if err != nil{
		fmt.Println("Open file err:",err)
		return
	}
	defer f.Close()
	//将字符串写入到文件
	//f.WriteString("helloword")
	//指定位置写入
	//f.WriteAt([]byte("china"),9)
	//字符切片写入
	//windows文本文件换行以\r\n为换行
	f.Write([]byte("I LOVE CHINA FOREVER"))
	fmt.Println("write successful")
	buf := make([]byte,1024)
	n,_ := f.ReadAt(buf,0)	 	//这里必须设置偏移量，因为上面写入后，光标停在了最后一个位置，而读取是从光标位置开始读取的
	fmt.Println(string(buf[:n]))
}

func Read_dict(){
	f,err := os.Open("E:/Go/dict.txt")
	if err!= nil{
		fmt.Println("Open Err:",err)
	}
	buf := make([]byte,1024)
	//读取整个文件的内容
	for  {
		n,err2 := f.Read(buf)
		if err2 != nil{
			fmt.Println("读取完毕")
			break
		}else{
			fmt.Println(string(buf[:n]))
		}
	}
}

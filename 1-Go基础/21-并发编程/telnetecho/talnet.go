package main

import (
	"fmt"
	"strings"
)

func processTelnetCommand(str string,exitChan chan int)bool{
	//@close指令表示终止本次会话
	if strings.HasPrefix(str,"@close"){
		fmt.Println("Session Closed")
		//告诉外部需要断开连接
		return false
	}else if strings.HasPrefix(str,"@shutdown"){
		fmt.Println("Server shudown")

		//往通道中写入0,阻塞等待接收方处理
		exitChan <- 0

		//告诉外部需要断开连接
		return false
	}
	//打印输入的字符串
	fmt.Println(str)
	return true
}
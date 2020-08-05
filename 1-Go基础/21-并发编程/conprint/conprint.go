package main

import "fmt"

func Printer (c chan int){

	//开始无限循环等待数据
	for {
		//从channel中获取一个数据
		data := <- c

		//将0 视为数据结束
		if data == 0 {
			break
		}
		fmt.Println("Printer接收到:",data)
	}

	//通知main 已经结束循环了
	fmt.Println("printer已结束")
	c <- 0		//阻塞 ，等待数据被接受

}
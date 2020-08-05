package main

import (
	"fmt"
)

func main() {
	//创建一个通道
	c := make(chan int)

	//并发执行printer 传入channel
	go Printer(c)

	for i := 1; i <= 10; i++ {
		//将数据通过channel投送给printer
		fmt.Printf("main --%d--> printer\n",i)
		c <- i
	}

	//通知并发的printer结束循环
	fmt.Println("没数据啦！")
	c <- 0

	//等待printer结束
	<- c
	fmt.Println("main结束")
}

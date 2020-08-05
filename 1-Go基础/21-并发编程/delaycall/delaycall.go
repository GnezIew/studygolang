package main

import (
	"fmt"
	"time"
)

func main() {
	//声明一个退出用的通道
	exit := make(chan int)

	//打印开始
	fmt.Println("start")

	//过1秒后调用匿名函数
	time.AfterFunc(time.Second, func() {

		//1秒后,打印结果
		fmt.Println("one second after")

		//通知main()的goroutine已经结束
		fmt.Println("匿名函数goroutine结束")
		exit <-0
	})

	//等待结束
	<- exit
	fmt.Println("main()的goroutine结束")
}

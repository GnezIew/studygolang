package main

import (
	"fmt"
	"time"
)

func running(){
	var times int
	//构建一个无限循环
	for {
		times ++
		fmt.Println("ticks",times)

		//延迟1秒
		time.Sleep(time.Second*2)
	}
}
func main() {

	//并发执行程序
	//go running()
	//
	////接受命令行输入,不做任何事情
	//var input string
	//fmt.Scanln(&input)

	//循环接受
	//构建一个通道
	//ch := make(chan int)
	//
	////开启一个并发匿名函数
	//go func() {
	//	for i:=3;i >=0  ; i-- {
	//		//发送数据
	//		fmt.Println("goroutine1 已发送数据:",i)
	//		ch <- i
	//
	//
	//		//每次发送完时等待
	//		time.Sleep(time.Second)
	//	}
	//}()
	//
	////遍历接受通道数据
	//for data := range ch {
	//	//打印通道数据
	//	fmt.Println("main goroutine 已接收数据:",data)
	//
	//	//当遇到0时，退出循环
	//	if data == 0{
	//		break
	//	}
	//}

	//--------------------带缓冲通道-----------------------//
	//创建一个3个元素缓冲大小的整型通道
	ch := make(chan int,3)

	//查看当前通道的大小
	fmt.Println(len(ch))
	ch <- 1
	ch <- 2
	ch <- 3
	//查看当前通道的大小
	fmt.Println(len(ch))

	data := <-ch
	fmt.Println(data)
	fmt.Println(len(ch))

	//通道不使用时可以调用进行关闭
	//close(ch)

}

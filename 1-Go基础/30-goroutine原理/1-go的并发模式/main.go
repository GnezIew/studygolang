package main

import (
	"fmt"
)

//生产者消费者模型

//生产者：生成factor整数倍的序列
func Producer(factor int,out chan<- int){
	for i:= 0; ;i++  {
		fmt.Println("生产了第",i*factor,"个包子")
		out <- i * factor
	}
}

//消费者
func Consumer(in <-chan int){
	for v:= range in{
		fmt.Println("吃了第",v,"个包子")
	}
}



func main() {
	////生产消费者模型
	//ch := make(chan int,64)  //成果队列
	//go Producer(3,ch)
	//go Producer(5,ch)
	//go Consumer(ch)
	//time.Sleep(5 * time.Second)

}

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	//序列号
	seq int64
)

//序列生成器
func GenID(i int) int64 {
	//尝试原子的增加序列号'
	fmt.Println("第",i,"个goroutine执行:")
	atomic.AddInt64(&seq,1)
	fmt.Println("第",i,"个goroutine执行结果:",seq)
	defer fmt.Println("第",i,"个goroutine执行结束")
	return seq
}
func main() {
	//生成10个并发序列号
	for i := 0;i < 10 ; i++  {
		go GenID(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(GenID(11))
}

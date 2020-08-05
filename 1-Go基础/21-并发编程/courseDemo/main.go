package main

import (
	"fmt"
	"sync"
)

func main() {
	//ch := make(chan struct{})
	//go func(){
	//	time.Sleep(time.Second)
	//	fmt.Println("协程完成")
	//	close(ch)
	//}()
	//fmt.Println("主函数执行")
	//<- ch		//从已关闭的通道中获取数据时，不会发生阻塞，如果没有数据，会接受该通道类型的零值，
	//fmt.Println("主函数执行完成")

	//------------------------------------------------
	//显示cpu核数
	//fmt.Println(runtime.GOMAXPROCS(0))

	//runtime.GOMAXPROCS(1)
	//intch := make(chan int, 1)
	//strch := make(chan string, 1)
	//intch <- 1
	//strch <- "hello"
	//select {
	////如果多个case都满足条件，会随机选择其中之一来执行
	//case value := <-intch:
	//	fmt.Println(value)
	//case value := <-strch:
	//	fmt.Println(value)
	//case <-time.After(time.Second * 5):
	//	fmt.Println("超时")
	//}
	//----------------------------
	var m sync.Mutex
	c := 0
	wg := sync.WaitGroup{}
	for i := 0;i < 100 ; i++  {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			fmt.Println(c)
			c++
		}()
	}
	wg.Wait()
	fmt.Println("完成",c)


}

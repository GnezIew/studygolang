package main

import (
	"fmt"
)

func add() {
	fmt.Println("函数")
}

//互斥锁
//var number int
//var mu sync.Mutex
//var wg sync.WaitGroup
//func test(){
//	mu.Lock()
//	defer func() {
//		mu.Unlock()	//解锁
//		wg.Done()
//	}()
//	number++
//}


func main() {
	//var pool sync.Pool
	//var val interface{}
	//pool.Put(add)
	//val = pool.Get()
	//v,ok := val.(func())
	//if ok{
	//	v()
	//}else{
	//	fmt.Println(1)
	//}

	//互斥锁
	//wg.Add(10000)
	//for i:=0;i< 10000 ;i++  {
	//	go test()
	//}
	//wg.Wait()
	//fmt.Println(number)
	//互斥锁 2
	//var mu sync.Mutex
	//var wg sync.WaitGroup
	//number := 0
	//wg.Add(10)
	//for i := 0;i < 10 ;i++  {
	//	go func(i int) {
	//		mu.Lock()
	//		fmt.Println("协程",i,"加锁")
	//		number++
	//		fmt.Println(number)
	//		defer func() {
	//			mu.Unlock()
	//			fmt.Println("协程",i,"解锁")
	//			wg.Done()
	//		}()
	//	}(i)
	//}
	//wg.Wait()

	//条件等待(sync.Cond)
	//var wg sync.WaitGroup
	//cond := sync.NewCond(new(sync.Mutex))
	//var number = 0
	//for i := 0; i < 3; i++ {
	//	go func(i int) {
	//		fmt.Println("协程",i,"启动")
	//		wg.Add(1)
	//		defer wg.Done()
	//		cond.L.Lock()
	//		fmt.Println("协程",i,"加锁")
	//		cond.Wait()
	//		fmt.Println("协程",i,"解锁")
	//		cond.L.Unlock()
	//		number++
	//		fmt.Println("协程",i,number)
	//	}(i)
	//}
	//time.Sleep(time.Second)
	//fmt.Println("主协程发送信号量1")
	//cond.Signal()
	//
	//time.Sleep(time.Second)
	//fmt.Println("主协程法索信号量2")
	//cond.Signal()
	//
	//time.Sleep(time.Second)
	//fmt.Println("主协程法索信号量3")
	//cond.Signal()
	//wg.Wait()

	//单次执行sync.Once
	//var once sync.Once
	//var wg sync.WaitGroup
	//
	//onceFunc := func() {
	//	fmt.Println("法师爱你们哟~")
	//}
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		defer wg.Done()
	//		once.Do(onceFunc)
	//	}()
	//}
	//wg.Wait()

	//并发安全Map(sync.Map)
	//var scence sync.Map
	////将键值对保存到sync.Map
	//scence.Store("法师",97)
	//scence.Store("老郑",100)
	//scence.Store("兵哥",200)
	//
	////从sync.Map中根据键取值
	//fmt.Println(scence.Load("法师"))
	////根据键删除对应的键值对
	//scence.Delete("法师")
	////遍历所有的sync.Map中的键值对
	//scence.Range(func(key, value interface{}) bool {
	//	fmt.Println("iterate:",key,":",value)
	//	return true
	//})


}

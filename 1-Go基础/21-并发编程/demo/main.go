package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	//逻辑中使用某个变量
	count int

	//与变量对应的使用互斥锁
	countGuard sync.RWMutex
)

func GetCount() int {
	//锁定
	countGuard.RLock()
	fmt.Println("GetCount()锁定了资源")

	//在函数退出时解除锁定
	defer func() {
		time.Sleep(time.Second *1)
		countGuard.RUnlock()
		fmt.Println("GetCount()解除了资源的锁定")
	}()
	return count
}

func GetCount2() int {
	//锁定
	countGuard.RLock()
	fmt.Println("GetCount()2锁定了资源")

	//在函数退出时解除锁定
	defer func() {
		countGuard.RUnlock()
		fmt.Println("GetCount()2解除了资源的锁定")
	}()
	return count
}

func SetCount(c int) {
	time.Sleep(time.Second * 2)
	countGuard.Lock()
	fmt.Println("SetCount()锁定了资源")
	count = c
	countGuard.Unlock()
	fmt.Println("SetCount()解除了资源的锁定")
}
func main() {

	go SetCount(1)

	go GetCount2()

	fmt.Println(GetCount())
	time.Sleep(time.Second*3)
}

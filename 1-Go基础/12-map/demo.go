package main

import (
	"fmt"
	"sync"
)

func demo(){
	var scene sync.Map

	//将键值对保存到sync.Map中
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)

	//从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	//从sync.Map中根据键删除对应的键值对
	scene.Delete("london")

	//遍历所有sync.Map中的键值对
	//只读已经存在的键
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:",key,value)
		return true
	})

	//go func() {
	//	for {
	//		scene.Store(1,1)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		fmt.Println(scene.Load(1))
	//	}
	//}()
	//
	//for{
	//	fmt.Println(2)
	//}
}

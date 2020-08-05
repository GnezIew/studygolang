package main

import (
	"fmt"
)

/*
闭包的最大用处有两个:
1.可以读取函数内部的变量
2.变量的值始终保持在内存中

应用场景:
1.网页访问量
2.
*/

func addDemo(a int, b int) int {
	return a + b
}

//将函数作为函数参数(函数回调)
func operator(f func(int, int) int) {
	value := f(10, 20)
	fmt.Println(value)
}

//将函数作为函数的返回值(闭包)
func getFunc()func ()int {
	i:=0
	return func() int {
		i+=1
		fmt.Printf("%p\n",&i)
		return i
	}
}

//闭包实现生成器
func playerGen(name string) func() (string,int){
	HP :=150
	return func() (string, int) {
		//将变量引用到闭包中
		return name,HP
	}
}
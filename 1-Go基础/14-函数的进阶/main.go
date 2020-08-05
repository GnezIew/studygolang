package main

import (
	"fmt"
)

//变量作用域
//全局变量
//定义全局变量
var num = 10
func testGlobal(){
	num := 100		//局部变量优先级高于全局变量
	//可以在函数中访问全局变量
	fmt.Println("全局变量",num)
}

func add(x,y int) int {
	return x + y
}
func sub(x,y int)int {
	return x-y
}

func multi(x,y int)int{
	return x * y
}
func calc(x,y int,op func(int,int) int) int {
	return op(x,y)
}

func main() {
	////外层不能访问到函数的内部变量(局部变量)
	//testGlobal()
	//for i := 0;i <=5;i++{
	//	fmt.Println(i)		//变量i只能 在for循环中被使用
	//}
	//---------------------------------------
	test := testGlobal		//函数可以作为变量
	fmt.Printf("%T\n",test)
	test()
	//-----------------------
	res := calc(100,50,add)
	fmt.Println(res)

	res2 := calc(100,200,sub)
	fmt.Println(res2)

	res3 := calc(100,200,multi)
	fmt.Println(res3)
}

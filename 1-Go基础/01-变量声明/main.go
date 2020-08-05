package main

import "fmt"

var m = 100	//全局变量

//定义有两个返回值的函数
func foo()(string,int){
	return "alex",9000
}

func main() {
	fmt.Println(m)
	//go语言声明变量必须要使用
	var name string
	var age int
	fmt.Println(name,age)

	//批量声明
	var (
		a string	//	没赋值默认 " "
		b int		//  默认 0
		c bool		//	默认为False
	)
	a = "沙河"
	b = 100
	c = true
	fmt.Println(a,b,c)

	//声明+初始化
	var x string = "曾为"
	fmt.Println(x)
	fmt.Printf("%s嘿嘿嘿\n",x)

	//类型推导(编译器根据变量初始值的类型指定给变量)
	var y = 200
	var z = true
	fmt.Println(y)
	fmt.Println(z)

	//简短变量声明，只能在函数内部使用
	m := 100	//局部变量
	fmt.Println(m)

	//匿名变量
	//调用foo函数
	// _ 用于接收不需要的变量值
	aa ,_ := foo()
	fmt.Println(aa)
	_,bb := foo()
	fmt.Println(bb)

	

	//在同一个作用域里不能重复声明同名变量
	var name2 = "nazha"
	var name3 = "zengwei"
	fmt.Println(name2,name3)


	//多重赋值
	res1,res2,res3 := 1,3,4
	res4,res5,res6 := 1,true,"嘿"
	fmt.Println(res1,res2,res3)
	fmt.Println(res4,res5,res6)

	//变量交换
	res1,res2 = res2, res1
	fmt.Println(res1,res2)
}

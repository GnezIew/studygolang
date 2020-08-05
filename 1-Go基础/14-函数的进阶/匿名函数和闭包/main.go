package main

import (
	"fmt"
	"strings"
)

//定义一个函数它的返回值是一个函数
//把函数作为了返回值
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

//闭包进阶实例:
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { //strings.HasSuffix :判断字符串后缀; 判断name是不是以suffix结尾
			return name + suffix
		}
		return name
	}
}

//闭包实例2
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

//匿名函数和闭包
func main() {
	sayhello := func() { //定义并赋值
		fmt.Println("匿名函数")
	}
	sayhello()

	func() {
		fmt.Println("匿名函数2")
	}() //定义并执行

	//闭包 = 函数 + 引用环境
	r := a("沙河哪吒") //闭包 = 函数 + 外层变量的引用，r此时就是一个闭包  //内部函数调用外部函数的变量
	r()            //相当于执行了a函数内部的匿名函数

	r2 := makeSuffixFunc(".txt")
	res := r2("沙河哪吒")
	fmt.Println(res)

	r3 := makeSuffixFunc(".avi")
	res2 := r3("沙河小王子")
	fmt.Println(res2)

	r4, r5 := calc(10)
	res3 := r4(10) //base = 10 + 10
	res4 := r5(10) //base = 20 - 10
	fmt.Println(res3)
	fmt.Println(res4)

	//打印函数的内存地址
	fmt.Println(addDemo)
	operator(addDemo)

	//闭包
	f := getFunc()
	fmt.Printf("%T\n", f)
	fmt.Println(f()) //i = 1
	fmt.Println(f()) //i = 2

	f1 := getFunc()
	fmt.Println(f1()) //i = 1	和上面的i是不同的变量

	/*
		循环退出时 i = 3
		defer 中是延迟调用的，所以结果为 3  3  3
	*/
	//for i := 0; i < 3; i++ {
	//	fmt.Println(&i)
	//	defer func() {
	//		fmt.Println(&i)
	//		fmt.Println(i)
	//	}()
	//}

	//---------------------------------------------------------
	//for i := 0; i < 3; i++ {
	//	fmt.Printf("for循环条件的i的地址:%p\n",&i)
	//	i:=i		//每次循环创建一个i的局部变量 (i的地址对不同)
	//	fmt.Printf("局部变量i的地址:%p\n",&i)
	//	defer func() {
	//		fmt.Printf("匿名函数中i的地址:%p\n",&i)
	//		fmt.Println(i)
	//	}()
	//}

	//---------------------------------------------------------
	//for j:=0;j<3 ;j++  {
	//	fmt.Printf("for循环的j的地址:%p\n",&j)
	//	defer func(j int) {				//函数参数值传递
	//		fmt.Printf("匿名函数参数j的地址:%p\n",&j)
	//		fmt.Println(j)
	//	}(j)
	//}

	//创建一个玩家生成器
	generator := playerGen("hight noon")

	//返回玩家的血量和姓名
	name , HP := generator()
	fmt.Println(name,HP)
}

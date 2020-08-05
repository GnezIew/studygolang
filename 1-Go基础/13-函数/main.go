package main

import "fmt"

//定义一个不带参数也没有返回值的函数
//func sayhello(){
//	fmt.Println("hello world")
//}

//定义一个接受string类型的name参数
//func sayhello2(name string){
//	fmt.Println("hello", name)
//}

//定义接受多个参数的函数
//func intSum(x int,y int) int {
//	ret := x + y
//	return ret
//}
//func intSum2(x int , y int) (ret int){
//	ret = x + y
//	return		//默认返回ret(可写可不写)
//}

//函数接受可变参数，在参数后面加... 表示可变参数
//可变参数在函数题中是切片类型
//固定参数和可变参数同时出现时，可变参数要放在最后
//go语言的函数中没有默认参数
//func intSum3(a int,b ... int) int{	//
//	fmt.Println(a)
//	fmt.Printf("%T\n",a)
//	Sum := 0
//	for _,v :=range b{
//		Sum = Sum + v
//	}
//	return Sum + a
//}

//Go语言中函数类型简写
//func intSum4(a,b int) int{		//如果参数类型一致，可以只写最后一个参数的类型
//	ret := a+b
//	return ret
//}

//支持多返回值
//func calc(a ,b int)(sum,sub int){	//多返回值也支持类型简写
//	sum = a + b		//因为在函数返回值那里已经定义了变量名，所以直接复制
//	sub = a - b
//	return sum,sub
//}

//-----------------------------------------------------------------
//defer语句 延迟执行




//主函数
func main() {
	//调用函数
	//sayhello()
	//sayhello2("李四")
	//name := "李四"
	//sayhello2(name)
	//res :=intSum(10,20)
	//fmt.Println(res)
	//
	//res2 := intSum3(10,20,30)
	//res3 := intSum3(10,20,30,40)
	//res4 := intSum3(10)	//a = 10,b =[]
	//fmt.Println(res2)
	//fmt.Println(res3)
	//fmt.Println(res4)

	//x,y := calc(100,200)
	//fmt.Println(x,y)

	//-------------------------
	//defer语句 延迟执行
	//在函数快要结束的时候执行defer的语句，按照逆序的方式执行(可以理解为先进后出)
	//defer语句常用来做：资源清理，文件关闭，解锁及记录时间等
	fmt.Println("start......")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end.......")

}

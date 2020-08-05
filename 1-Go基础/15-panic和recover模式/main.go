package main

/*
	1.编辑时异常
	2.编译时异常
	3.运行时异常
*/


import "fmt"

//panic 和 recover
//func a()  {
//	fmt.Println("func a")
//}
//func b()  {
//	defer func() {
//		err := recover()		//错误收集
//		if err != nil{
//			fmt.Println("func b in error")
//		}else{
//
//		}
//	}()
//	panic("panic in b")		//抛出异常
//}
//func c()  {
//	fmt.Println("func c")
//}

func demo(i int){
	var arr [10]int
	//设置错误拦截 defer + 匿名函数 +recover
	defer func() {
		err := recover()	//错误拦截
		if err != nil{
			fmt.Println(err)
		}
	}()
	//数组下标越界
	arr[i] =123
	//出现错误后续代码不会继续执行
	/*
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	*/
}

func main() {
	//a()
	//b()
	//c()
	//---------------------------------
	//panic错误
	//defer func(){
	//	//	recover 捕获错误
	//	//recover一定要搭配defer使用
	//	//recover一定放字panic之前定义
	//	err := recover()	//尝试将函数从当前的异常状态恢复过来
	//	fmt.Println("recover捕获到了panic异常",err)
	//}()
	//var a []int
	//a[0] = 100

	demo(10)
	fmt.Println("程序继续执行")
}

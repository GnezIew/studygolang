package main

//Go和C 混合编程 需要传递参数和接受返回值

/*
#include <stdio.h>
int add(int a,int b)
{
	return a+b;
}
*/
import "C"
import "fmt"

func main() {
	a:=10
	b:=20
	//将Go语言的类型转换为C语言类型
	c:=C.add(C.int(a),C.int(b))
	fmt.Println(c)
	fmt.Printf("%T\n",c)			//main._Ctype_int
	fmt.Println(c+C.int(b))
}

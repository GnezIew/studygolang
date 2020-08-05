package main

import (
	"fmt"
	"math"
)

//常量是不允许寻址
const  pi  = 3.14	//全局常量
//常量是不允许被修改的

//批量声明常量
const (
	a = 100
	b = 50
	c
	d =30
	e			//默认和上一个的值一样
)

//const (
//	n1 = iota
//	n2
//	_
//	n3
//)
const (
	n1 = iota
	n2 = 100
	n3
	n4 = iota
)
func main() {
	fmt.Println(pi)
	fmt.Println(a,b,c,d,e)

	//iota 常量计数器，枚举;const中每增加一行iota就加1
	const (
		aa = iota	//默认在const出现时重置为零
		bb
		cc
		dd
	)
	fmt.Println(aa,bb,cc,dd)
	fmt.Println(n1,n2,n3,n4)

	//常量一般建议大写
	const PI = 3.1415926
	fmt.Println(PI)
	fmt.Println(math.Pi)

}